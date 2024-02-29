package bands

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oracao-bandas.com/src/api/modules/bands/controllers"
	"oracao-bandas.com/src/api/modules/bands/services"
	"oracao-bandas.com/src/database/resources/bands/repositories"
)

func InitBandModule(group *gin.RouterGroup, database *gorm.DB) {
	bandRepository := repositories.NewPostgresBandRepository(database)
	bandService := services.NewSaveBandService(bandRepository)
	saveBandController := controllers.NewSaveBandController(bandService)
	SetupRoutes(group, saveBandController)
}
