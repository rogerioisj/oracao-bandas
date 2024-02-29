package bands

import (
	"github.com/gin-gonic/gin"
	"oracao-bandas.com/src/api/modules/bands/controllers"
)

func SetupRoutes(router *gin.RouterGroup, controller controllers.SaveBandControllerInterface) {
	bands := router.Group("bands")

	bands.POST("", controller.Save)
}
