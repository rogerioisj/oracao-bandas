package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oracao-bandas.com/src/modules/auth/services"
	auth "oracao-bandas.com/src/modules/auth/structs"
)

var (
	sessions   = make(map[string]string)
	cookieName = "session_token"
)

type AuthControllerInterface interface {
	RegisterUser(context *gin.Context)
	Login(context *gin.Context)
}

type AuthController struct {
	service services.AuthServiceInterface
}

func NewAuthController(service services.AuthServiceInterface) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (controller *AuthController) RegisterUser(context *gin.Context) {
	var dto auth.CreateUserDto

	login := context.PostForm("login")
	password := context.PostForm("password")
	name := context.PostForm("name")

	if login == "" || password == "" || name == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "No empty values allowed",
		})

		return
	}

	dto.Name = name
	dto.Login = login
	dto.Password = password

	err := controller.service.CreateUser(&dto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Error to create user",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{})

	return
}

func (controller *AuthController) Login(context *gin.Context) {
	var dto auth.LoginDto

	login := context.PostForm("login")
	password := context.PostForm("password")

	if login == "" || password == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "No empty values allowed",
		})

		return
	}

	dto.Login = login
	dto.Password = password

	err := controller.service.Login(&dto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Error to create user",
		})
		return
	}

	sessionToken := "session_token" // generate a unique session token, you may want to use UUID or JWT
	sessions[sessionToken] = dto.Login

	context.SetCookie(cookieName, sessionToken, 3600, "/", "", false, true)

	context.JSON(http.StatusOK, gin.H{})
}
