package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/michaelbui/bingo-backend/models"
)

type AdminController struct {
}

var (
	adminController *AdminController
)

func Admin() *AdminController {
	if adminController == nil {
		adminController = &AdminController{}
	}
	return adminController
}

func (controller *AdminController) Reset(context echo.Context) error {
	if err := models.Game().Reset(); err != nil {
		return context.JSON(http.StatusInternalServerError, JsonResponse{
			Code:    1,
			Message: "Error While Resetting The Game!",
		})
	}
	return context.JSON(http.StatusOK, JsonResponse{Message: "The Game Has Been Reset!"})
}

func (controller *AdminController) Activate(context echo.Context) error {
	if err := models.Game().Activate(); err != nil {
		return context.JSON(http.StatusInternalServerError, JsonResponse{
			Code:    1,
			Message: "Error While Activating The Game!",
		})
	}
	return context.JSON(http.StatusOK, JsonResponse{Message: "The Game Has Been Activated!"})
}
