package bands

import (
	"github.com/gin-gonic/gin"
	"oracao-bandas.com/src/modules/bands/controllers"
)

func SetupRoutes(router *gin.RouterGroup, controller controllers.SaveBandControllerInterface) {
	bands := router.Group("bands")

	bands.POST("", controller.Save)
}
