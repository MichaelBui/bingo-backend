package controllers

import (
	"github.com/labstack/echo"
	"github.com/michaelbui/bingo-backend/models"
	"net/http"
)

type NumberController struct {

}

var (
	numberController *NumberController
)

func Number() *NumberController {
	if numberController == nil {
		numberController = &NumberController{}
	}
	return numberController
}

func (n *NumberController) List(context echo.Context) error {
	timestamps, values := models.Number().List()
	return context.JSON(http.StatusOK, JsonResponse{Data: map[string]interface{}{
		"timestamps": timestamps,
		"values": values,
	}})
}

func (n *NumberController) Next(context echo.Context) error {
	if !models.Game().IsLocked() {
		return context.JSON(http.StatusBadRequest, JsonResponse{Code: 2, Message: "Game Not Started"})
	}
	number := models.Number().Next()
	if number == 0 {
		return context.JSON(http.StatusBadRequest, JsonResponse{Code: 1, Data: number})
	}
	return context.JSON(http.StatusOK, JsonResponse{Data: number})
}