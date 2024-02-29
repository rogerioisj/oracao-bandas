package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Functional server",
		"data":    time.Now(),
	})
}

func SetupRoutes(router *gin.RouterGroup) {
	home := router.Group("home")

	home.GET("/health", HealthCheck)
}
