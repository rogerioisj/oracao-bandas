package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oracao-bandas.com/src/modules/home/controllers"
	"time"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Functional server",
		"data":    time.Now(),
	})
}

func SetupRoutes(router *gin.RouterGroup, controller controllers.HomeControllerInterface) {
	router.GET("/home/health", HealthCheck)
	router.GET("", controller.Home)
}
