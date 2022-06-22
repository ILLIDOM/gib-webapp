package http

import (
	"net/http"

	"github.com/ILLIDOM/gin-webapp/cmd/authentication"
	"github.com/ILLIDOM/gin-webapp/cmd/domain"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

func (loginHandler *LoginHandler) Login(context *gin.Context) {
	var loginObject authentication.LoginRequest
	var user domain.User
	if err := context.ShouldBindJSON(&loginObject); err != nil {
		// TODO: add custom error message
		context.AbortWithStatusJSON(http.StatusBadRequest, nil)
	}

	//TODO: validate provided username and password

	//TODO: load user from db

	tokenString, err := authentication.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
