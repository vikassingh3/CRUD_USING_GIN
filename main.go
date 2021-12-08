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
	r.GET("/:id", handler.GetUser)
	r.DELETE("/delete/:id", handler.DeleteUser)
	r.PUT("/update/:id", handler.UpdateUser)
	r.DELETE("/delete", handler.DeleteAll)

	godotenv.Load()

	database.ConnectDB()
	r.Run(":8080")

}
