package http

import (
	"net/http"

	"github.com/ILLIDOM/gin-webapp/cmd/authentication"
	"github.com/ILLIDOM/gin-webapp/cmd/database"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	UserService database.UserService
}

func NewLoginHandler(userService database.UserService) *LoginHandler {
	return &LoginHandler{
		UserService: userService,
	}
}

func (loginHandler *LoginHandler) Login(context *gin.Context) {
	var loginObject authentication.LoginRequest
	if err := context.ShouldBindJSON(&loginObject); err != nil {
		// TODO: add custom error message
		context.AbortWithStatusJSON(http.StatusBadRequest, nil)
	}

	user, err := loginHandler.UserService.GetByID(loginObject.UserID)
	if err != nil {
		// no user found with the provided UserID
		context.AbortWithStatus(http.StatusUnauthorized)
	}
	// validate provided username and password - compare login request with user
	if user.Fullname != loginObject.UserName || user.Password != loginObject.Password {
		context.AbortWithStatus(http.StatusUnauthorized)
	}

	user.Roles, err = loginHandler.UserService.GetRolesByUserID(user.ID)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	}

	tokenString, err := authentication.GenerateJWT(*user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
