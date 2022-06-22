package middleware

import "github.com/gin-gonic/gin"

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement
	}
}

func Authorization(validRoles []int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement
	}
}
