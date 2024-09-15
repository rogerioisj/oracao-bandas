package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"oracao-bandas.com/src/modules/auth/services"
	auth "oracao-bandas.com/src/modules/auth/structs"
	"strconv"
)

var (
	sessions   = make(map[string]string)
	cookieName = "session_token"
)

type AuthControllerInterface interface {
	RegisterUser(context *gin.Context)
	Login(context *gin.Context)
	Logout(context *gin.Context)
	ListUsers(context *gin.Context)
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

func (controller *AuthController) ListUsers(context *gin.Context) {
	var maxPage int

	query := context.Request.URL.Query()

	page, _ := strconv.Atoi(query.Get("page"))
	itens, _ := strconv.Atoi(query.Get("itens"))
	name := query.Get("name")

	if page <= 0 {
		page = 1
	}

	if itens <= 0 {
		itens = 5
	}

	users, total := controller.service.ListUsers(page, itens, name)

	maxPage = total / itens

	i := total % itens
	if i > 0 {
		maxPage += 1
	}

	_, err := context.Cookie("session_token")
	cookieExists := err == nil

	log.Printf("Itens: %v", itens)
	log.Printf("Total itens: %v", total)
	log.Printf("Max page: %v", maxPage)
	log.Printf("Page: %v", page)
	log.Printf("Cookie: %v", cookieExists)

	for _, v := range users {
		log.Printf("Name: %s; User: %s, Created: %s, Email: %s", v.Name, v.Login, v.CreatedAt)
	}

	context.HTML(http.StatusOK, "users.html", gin.H{
		"users":        &users,
		"page":         &page,
		"itens":        &itens,
		"cookieExists": &cookieExists,
		"maxPage":      &maxPage,
	},
	)
}
