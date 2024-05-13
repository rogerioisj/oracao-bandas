package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oracao-bandas.com/src/database/repositories"
	"oracao-bandas.com/src/modules/auth/controllers"
	"oracao-bandas.com/src/modules/auth/services"
)

func InitAuthModule(group *gin.RouterGroup, database *gorm.DB) {
	repository := repositories.NewUserRepository(database)
	service := services.NewAuthService(repository)
	controller := controllers.NewAuthController(service)
	SetupRoutes(group, controller)
}
