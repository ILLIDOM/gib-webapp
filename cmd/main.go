package main

import (
	"database/sql"
	"fmt"

	"github.com/ILLIDOM/gin-webapp/cmd/database"
	"github.com/ILLIDOM/gin-webapp/cmd/http"
	"github.com/ILLIDOM/gin-webapp/cmd/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	//create db connection
	db, err := sql.Open("sqlite3", "../database.db")
	if err != nil {
		fmt.Printf("Can't open db: %v", err)
	}

	// create User service with reference to db
	userService := database.UserService{
		DB: db,
	}
	//create handlers
	loginHandler := http.NewLoginHandler(userService)
	userHandler := http.NewHandler(userService)

	// create http server without any middleware
	server := gin.New()
	// attach middleware to server - smae as gin.Default()
	server.Use(gin.Recovery())
	server.Use(gin.Logger())

	server.POST("/login", loginHandler.Login)

	api := server.Group("/api")
	api.Use(middleware.ValidateToken())

	admin := api.Group("/admin")
	admin.Use(middleware.Authorization([]int{1}))

	server.GET("/users/:user_id", userHandler.GetByID)
	server.POST("/users", userHandler.Create)

	server.Run(":8080")
}
