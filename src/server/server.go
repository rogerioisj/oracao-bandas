package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"html/template"
	"oracao-bandas.com/src/configuration"
	"oracao-bandas.com/src/database/resources/bands/entities"
	"oracao-bandas.com/src/modules"
)

func StartServer(config *configuration.Configuration, db *gorm.DB) {
	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"isEven": func(value int) bool {
			return value%2 == 0
		},
		"greaterThan": func(array []entities.Band, value int) bool { return len(array) > value },
	})

	router.LoadHTMLGlob("src/views/**/*")
	router.Static("/static", "src/static/")

	api := router.Group("/")

	modules.InitModules(api, db)

	err := router.Run(config.System.Port)
	if err != nil {
		return
	}
}
