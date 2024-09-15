package auth

import (
	"github.com/gin-gonic/gin"
	"oracao-bandas.com/src/middlewares"
	"oracao-bandas.com/src/modules/auth/controllers"
)

func SetupRoutes(router *gin.RouterGroup, controller controllers.AuthControllerInterface, middlewares middlewares.SessionCheckInterface) {
	authGroup := router.Group("auth")

	authGroup.POST("register-user", middlewares.CheckSessionMiddleware(), controller.RegisterUser)
	authGroup.POST("login", controller.Login)
	authGroup.GET("logout", controller.Logout)
	authGroup.GET("users", controller.ListUsers)
	authGroup.POST("update-user", controller.UpdateLoggedUser)
}
