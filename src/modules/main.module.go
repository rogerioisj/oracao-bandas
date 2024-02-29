package modules

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oracao-bandas.com/src/modules/bands"
	"oracao-bandas.com/src/modules/home"
)

func InitModules(api *gin.RouterGroup, database *gorm.DB) {
	bands.InitBandModule(api, database)
	home.SetupRoutes(api)
}
