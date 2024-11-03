package middleware

import (
	auth "go-4-digi-micro/auth/proto"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthMiddleware(authClient auth.AuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// prepare payload to auth service
		req := &auth.AuthValidateRequest{
			Token: tokenString,
		}

		// call auth service - Validate
		res, err := authClient.Validate(c, req)
		if err != nil {
			if e, ok := status.FromError(err); ok {
				switch e.Code() {
				case codes.Unauthenticated:
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": e.Message()})
				default:
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": e.Message()})
				}
			} else {

			}
			return
		}

		// Set the token claims to the context
		c.Set("auth_id", res.AuthId)
		c.Set("account_id", res.AccountId)
		c.Set("username", res.Username)

		c.Next() // Authorized, Proceed to the next handler
	}
}
