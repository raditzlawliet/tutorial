package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	AuthLogin(*gin.Context)
}

type authImplement struct{}

func NewAuth() AuthInterface {
	return &authImplement{}
}

type authLoginPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *authImplement) AuthLogin(c *gin.Context) {
	bodyPayloadAuth := authLoginPayload{}

	err := c.ShouldBindJSON(&bodyPayloadAuth)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("`%v` Login Sukses", bodyPayloadAuth.Username),
		"data":    nil,
	})
}
