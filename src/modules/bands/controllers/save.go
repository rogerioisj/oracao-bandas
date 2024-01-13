package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oracao-bandas.com/src/modules/bands/services"
	"oracao-bandas.com/src/modules/bands/structs"
)

func Save(context *gin.Context) {
	var band structs.SaveBandDto

	dtoError := context.ShouldBindJSON(&band)

	if dtoError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": dtoError.Error(),
		})
	}

	serviceError, serviceReturn := services.SaveBand(&band)

	if serviceError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": serviceError.Error(),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"band": serviceReturn,
	})
}
