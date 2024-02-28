package server

import (
	"github.com/gin-gonic/gin"
	"oracao-bandas.com/src/api/modules/bands"
	"oracao-bandas.com/src/api/modules/home"
	"oracao-bandas.com/src/controllers"
)

func StartServer() {
	router := gin.Default()

	router.LoadHTMLGlob("src/views/**/*")
	router.Static("/static", "src/static/")

	api := router.Group("api")

	home.SetupRoutes(api)
	bands.SetupRoutes(api)

	router.GET("/", controllers.Home)

	err := router.Run(":8000")
	if err != nil {
		return
	}
}
