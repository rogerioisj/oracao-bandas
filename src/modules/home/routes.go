package home

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Functional server",
		"data":    time.Now(),
	})
}

func postTest(c *gin.Context) {
	var test Book

	err := c.ShouldBindJSON(&test)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	fmt.Printf("Title: %s \nAuthor: %s \n", test.Title, test.Author)
}

func SetupRoutes(router *gin.RouterGroup) {
	home := router.Group("home")

	home.GET("/health", HealthCheck)
	home.POST("/test", postTest)
}
