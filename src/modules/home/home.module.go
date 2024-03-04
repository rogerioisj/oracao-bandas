package home

import (
	"github.com/gin-gonic/gin"
	"oracao-bandas.com/src/modules/bands/services"
	"oracao-bandas.com/src/modules/home/controllers"
)

func InitHomeModule(group *gin.RouterGroup, bandService *services.SaveBandService) {
	saveBandController := controllers.NewHomeController(bandService)
	SetupRoutes(group, saveBandController)
}
