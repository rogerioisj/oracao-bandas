package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oracao-bandas.com/src/modules/bands/services"
)

type Message struct {
	Text string
}

type HomeControllerInterface interface {
	Home(context *gin.Context)
}

type HomeController struct {
	service services.SaveBandServiceInterface
}

func NewHomeController(service services.SaveBandServiceInterface) *HomeController {
	return &HomeController{
		service: service,
	}
}

func (h HomeController) Home(ctx *gin.Context) {
	lastBands := h.service.LoadLastBands()

	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"lastBands":  &lastBands,
		"queryBands": &lastBands,
	},
	)
}
