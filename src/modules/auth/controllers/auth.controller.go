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

	context.Redirect(http.StatusFound, "/")

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

	sessionData, err := controller.service.Login(&dto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Error to create user",
		})
		return
	}

	sessionToken := sessionData
	sessions[sessionToken] = dto.Login

	context.SetCookie(cookieName, sessionToken, 36000, "/", "", false, true)

	context.Redirect(http.StatusFound, "/")
}

func (controller *AuthController) Logout(context *gin.Context) {
	sessionToken, err := context.Cookie(cookieName)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "No session found"})
		return
	}

	controller.service.Logout(sessionToken)

	// Clear cookie
	context.SetCookie(cookieName, "", -1, "/", "", false, true)

	context.Redirect(http.StatusFound, "/")
}
