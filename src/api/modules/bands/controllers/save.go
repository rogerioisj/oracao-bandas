package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oracao-bandas.com/src/api/modules/bands/services"
	"oracao-bandas.com/src/api/modules/bands/structs"
)

type SaveBandControllerInterface interface {
	Save(context *gin.Context)
}

type SaveBandController struct {
	service services.SaveBandServiceInterface
}

func NewSaveBandController(service services.SaveBandServiceInterface) *SaveBandController {
	return &SaveBandController{
		service: service,
	}
}

func (controller *SaveBandController) Save(context *gin.Context) {
	var band structs.SaveBandDto

	dtoError := context.ShouldBindJSON(&band)

	if dtoError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": dtoError.Error(),
		})
	}

	serviceError, serviceReturn := controller.service.SaveBand(&band)

	if serviceError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": serviceError.Error(),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"band": serviceReturn,
	})
}
