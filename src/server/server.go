package server

import (
	"github.com/gin-gonic/gin"
	"oracao-bandas.com/src/api/modules/bands"
	"oracao-bandas.com/src/api/modules/home"
	"oracao-bandas.com/src/configuration"
	"oracao-bandas.com/src/controllers"
)

func StartServer(config *configuration.Configuration) {
	router := gin.Default()

	router.LoadHTMLGlob("src/views/**/*")
	router.Static("/static", "src/static/")

	api := router.Group("api")

	home.SetupRoutes(api)
	bands.SetupRoutes(api)

	router.GET("/", controllers.Home)

	err := router.Run(config.System.Port)
	if err != nil {
		return
	}
}
