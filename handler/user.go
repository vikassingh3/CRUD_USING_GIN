package handler

import (
	"net/http"
	"path/database"
	"path/modelss"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := database.DB

	var user = new(modelss.User)

	c.BindJSON(&user)

	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "user created",
		"data": user,
	})
}

func GetAllUser(c *gin.Context) {

	db := database.DB
	var users []modelss.User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "get all users",
		"data": users,
	})

}
func GetUser(c *gin.Context) {
	id := c.Param("id")

	db := database.DB
	var user modelss.User

	db.Find(&user, id)

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no data in the users collection",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"err":  false,
		"msg":  "found the data with this id",
		"user": user,
	})
	return
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	db := database.DB

	var user modelss.User

	db.Find(&user, id).Delete(&user, id)

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "false",
			"msg":    "no product wiht this id",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"msg":    "delete the user ",
	})
	return

}

func DeleteAll(c *gin.Context) {
	db := database.DB

	var users modelss.User
	c.BindJSON(&users)

	if users.Names == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "no data found",
		})
		return
	}

	db.Delete(&users)
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"msg":    "all users deleted successfully",
	})

}

func UpdateUser(c *gin.Context) {
	var s = new(modelss.User)
	c.BindJSON(&s)

	id := c.Param("id")

	db := database.DB

	var data = new(modelss.User)

	db.Find(&data, id)

	data.Names = s.Names
	data.Username = s.Username
	data.Email = s.Email

	if s.Names == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "data not found witht this ID",
		})

	}

	db.Update(&data)
	db.Save(&data)

	c.JSON(http.StatusOK, gin.H{
		"message": "data updated  sucessfull",
		"data":    data,
	})

	return

}
