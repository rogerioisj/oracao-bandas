package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct {
	Text string
}

func Home(ctx *gin.Context) {
	data := Message{
		Text: "Ronaldo!",
	}

	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"Message": data,
	})
}
