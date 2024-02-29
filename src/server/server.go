package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"html/template"
	"oracao-bandas.com/src/api/modules"
	"oracao-bandas.com/src/configuration"
	"oracao-bandas.com/src/controllers"
)

func StartServer(config *configuration.Configuration, db *gorm.DB) {
	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"isEven": func(value int) bool {
			return value%2 == 0
		},
	})

	router.LoadHTMLGlob("src/views/**/*")
	router.Static("/static", "src/static/")

	api := router.Group("api")

	modules.InitModules(api, db)

	router.GET("/", controllers.Home)

	err := router.Run(config.System.Port)
	if err != nil {
		return
	}
}
