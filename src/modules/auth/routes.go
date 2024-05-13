package auth

import (
	"github.com/gin-gonic/gin"
	"oracao-bandas.com/src/modules/auth/controllers"
)

func SetupRoutes(router *gin.RouterGroup, controller controllers.AuthControllerInterface) {
	authGroup := router.Group("auth")

	authGroup.POST("register-user", controller.RegisterUser)
	authGroup.POST("login", controller.Login)
	authGroup.GET("logout", controller.Logout)
}
