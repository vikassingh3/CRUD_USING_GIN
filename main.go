package main

import (
	// "fmt"
	"path/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"path/database"
)

func main() {
	r := gin.Default()
	r.POST("/create", handler.CreateUser)
	r.GET("/get", handler.GetAllUser)

	godotenv.Load()

	database.ConnectDB()
	r.Run(":8080")

}

/*
func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("err while loading .env")
	}

}
*/
