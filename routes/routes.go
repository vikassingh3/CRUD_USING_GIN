package routes

import (
	"path/handler"

	"github.com/gin-gonic/gin"
)

func Router(app *gin.Context) {
	r := gin.Default()

	r.GET("/carete", handler.CreateUser)

}
