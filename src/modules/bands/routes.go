package bands

import (
	"github.com/gin-gonic/gin"
	"oracao-bandas.com/src/modules/bands/controllers"
)

func SetupRoutes(router *gin.RouterGroup) {
	bands := router.Group("bands")

	bands.POST("", controllers.Save)
}
