package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ILLIDOM/gin-webapp/cmd/database"
	"github.com/ILLIDOM/gin-webapp/cmd/http"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	// create db
	file, err := os.Create("database.db")
	if err != nil {
		fmt.Println(err)
	}
	file.Close()

	//create db connection
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Printf("Can't open db: %v", err)
	}

	//populate db
	createTable(db)

	// create User service with reference to db
	userService := database.UserService{
		DB: db,
	}
	//create httpHandler for Users
	userHandler := http.NewHandler(userService)

	router.GET("/users/:user_id", userHandler.GetByID)
	router.POST("/users", userHandler.Create)

	router.Run(":8080")
}

func createTable(db *sql.DB) {
	users_table := `CREATE TABLE users (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "Firstname" TEXT,
        "Lastname" TEXT,
        "Fullname" TEXT,
        "Email" TEXT);`
	query, err := db.Prepare(users_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()
	fmt.Println("Table created successfully!")
}
