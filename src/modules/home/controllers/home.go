package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oracao-bandas.com/src/database/resources/bands/entities"
	"oracao-bandas.com/src/modules/bands/services"
	"strconv"
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
	var loadedBands []entities.Band

	query := ctx.Request.URL.Query()

	page, _ := strconv.Atoi(query.Get("page"))
	itens, _ := strconv.Atoi(query.Get("itens"))

	if page <= 0 {
		page = 1
	}

	if itens <= 0 {
		itens = 5
	}

	loadedBands = h.service.SearchBands(page, itens)

	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"lastBands":  &lastBands,
		"queryBands": &loadedBands,
	},
	)
}
