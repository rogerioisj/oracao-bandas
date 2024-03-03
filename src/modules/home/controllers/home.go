package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
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
	var total int
	var maxPage int

	query := ctx.Request.URL.Query()

	page, _ := strconv.Atoi(query.Get("page"))
	itens, _ := strconv.Atoi(query.Get("itens"))

	if page <= 0 {
		page = 1
	}

	if itens <= 0 {
		itens = 5
	}

	loadedBands, total = h.service.SearchBands(page, itens)

	maxPage = total / itens

	i := total % itens
	if i > 0 {
		maxPage += 1
	}

	log.Printf("Itens: %v", total)
	log.Printf("Total itens: %v", total)
	log.Printf("Max page: %v", maxPage)
	log.Printf("Page: %v", page)

	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"lastBands":  &lastBands,
		"queryBands": &loadedBands,
		"page":       &page,
		"itens":      &itens,
		"maxPage":    &maxPage,
	},
	)
}
