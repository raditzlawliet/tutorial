package handler

import (
	"fmt"
	"go-4-digi-micro/account/model"
	auth "go-4-digi-micro/auth/proto"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AuthInterface interface {
	Login(*gin.Context)
	Upsert(*gin.Context)
}

type authImplement struct {
	authClient auth.AuthClient
	db         *gorm.DB
}

func NewAuth(authClient auth.AuthClient, db *gorm.DB) AuthInterface {
	return &authImplement{
		authClient,
		db,
	}
}

type authLoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *authImplement) Login(c *gin.Context) {
	payload := authLoginPayload{}

	// parsing payload to struct
	err := c.BindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// prepare payload to auth service
	req := &auth.AuthLoginRequest{
		Username: payload.Username,
		Password: payload.Password,
	}

	// call auth service Login
	res, err := a.authClient.Login(c, req)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Unauthenticated:
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": e.Message()})
			case codes.Internal:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": e.Message()})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": e.Message()})
			}
		}
		return
	}

	// response to client
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v Login Sukses", payload.Username),
		"data":    res.Token,
	})
}

type authUpsertPayload struct {
	AccountID int64  `json:"account_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (a *authImplement) Upsert(c *gin.Context) {
	payload := authUpsertPayload{}

	// parsing payload to struct
	err := c.BindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// Check AccountID is valid
	var account model.Account
	if err := a.db.First(&account, payload.AccountID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Account Not found",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// prepare payload to auth service
	req := &auth.AuthUpsertRequest{
		AccountId: payload.AccountID,
		Username:  payload.Username,
		Password:  payload.Password,
	}

	// call auth service Upsert
	_, err = a.authClient.Upsert(c, req)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e.Message()})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": e.Message()})
			}
		}
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Create success",
		"data":    payload.Username,
	})
}
