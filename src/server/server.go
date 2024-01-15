package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oracao-bandas.com/src/modules/bands"
	"oracao-bandas.com/src/modules/home"
)

func StartServer() {
	router := gin.Default()

	api := router.Group("api")

	home.SetupRoutes(api)
	bands.SetupRoutes(api)

	api.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	err := router.Run(":8000")
	if err != nil {
		return
	}
}
