package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oracao-bandas.com/src/modules/home"
)

func StartServer() {
	router := gin.Default()

	api := router.Group("api")

	home.SetupRoutes(api)

	api.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	router.Run()
}
