package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"html/template"
	"oracao-bandas.com/src/configuration"
	"oracao-bandas.com/src/database/entities"
	"oracao-bandas.com/src/modules"
)

func StartServer(config *configuration.Configuration, db *gorm.DB) {
	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"isEven": func(value int) bool {
			return value%2 == 0
		},
		"greaterThanArray": func(array []entities.Band, value int) bool { return len(array) > value },
		"minusOrEqualThan": func(value1, value2 int) bool { return value1 <= value2 },
		"greaterThan":      func(value1, value2 int) bool { return value1 > value2 },
		"minusThan":        func(value1, value2 int) bool { return value1 < value2 },
		"sub":              func(value1, value2 int) int { return value1 - value2 },
		"add":              func(value1, value2 int) int { return value1 + value2 },
		"equalsBool":       func(value1, value2 bool) bool { return value1 == value2 },
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
