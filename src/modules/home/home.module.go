package home

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oracao-bandas.com/src/database/resources/bands/repositories"
	"oracao-bandas.com/src/modules/bands/services"
	"oracao-bandas.com/src/modules/home/controllers"
)

func InitHomeModule(group *gin.RouterGroup, database *gorm.DB) {
	bandRepository := repositories.NewPostgresBandRepository(database)
	bandService := services.NewSaveBandService(bandRepository)
	saveBandController := controllers.NewHomeController(bandService)
	SetupRoutes(group, saveBandController)
}
