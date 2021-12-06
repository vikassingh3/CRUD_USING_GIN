package handler

import (
	"net/http"
	"path/database"
	"path/modelss"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := database.DB

	user := new(modelss.User)

	c.BindJSON(&user)

	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "user created",
		"data": user,
	})

}

func GetAllUser(c *gin.Context) {

	db := database.DB
	users := new([]modelss.User)
	c.BindJSON(&users)
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "get all users",
		"data": users,
	})

}
