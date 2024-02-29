package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct {
	Text string
}

func Home(ctx *gin.Context) {
	data := [...]Message{
		Message{
			Text: "Ronaldo!",
		},
		Message{
			Text: "Brilha muito no curintinhans!",
		},
	}

	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"data": data,
	},
	)
}
