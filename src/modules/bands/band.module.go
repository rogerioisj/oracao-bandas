package bands

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oracao-bandas.com/src/database/resources/bands/repositories"
	"oracao-bandas.com/src/modules/bands/controllers"
	"oracao-bandas.com/src/modules/bands/services"
)

func InitBandModule(group *gin.RouterGroup, database *gorm.DB) *services.SaveBandService {
	bandRepository := repositories.NewPostgresBandRepository(database)
	bandService := services.NewSaveBandService(bandRepository)
	saveBandController := controllers.NewSaveBandController(bandService)
	SetupRoutes(group, saveBandController)

	return bandService
}
