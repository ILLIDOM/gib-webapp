package http

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ILLIDOM/gin-webapp/cmd/database"
	"github.com/ILLIDOM/gin-webapp/cmd/domain"
	"github.com/ILLIDOM/gin-webapp/cmd/utils/errors"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService database.UserService
}

func NewHandler(userService database.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (userHandler *UserHandler) GetByID(ctx *gin.Context) {
	userID, err := strconv.Atoi(strings.TrimSpace(ctx.Param("user_id")))
	if err != nil {
		fmt.Printf("Error parsing userID: %v", err)
	}
	user, err := userHandler.UserService.GetByID(userID)
	if err != nil {
		restErr := errors.NewRestError("user not found", 404, err.Error())
		ctx.JSON(restErr.ErrStatus, restErr)
	}
	ctx.JSON(http.StatusOK, user)
}

func (userHandler *UserHandler) Create(ctx *gin.Context) {
	var user domain.User
	if err := ctx.BindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.ErrStatus, restErr)
	}
	if _, err := userHandler.UserService.Create(user); err != nil {
		restErr := errors.NewBadRequestError("cant create resource")
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}
	ctx.JSON(http.StatusCreated, user)
}
