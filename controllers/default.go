package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/michaelbui/bingo-backend/helpers"
)

type DefaultController struct {

}

var (
	defaultController *DefaultController
)

func Default() *DefaultController {
	if defaultController == nil {
		defaultController = &DefaultController{}
	}
	return defaultController
}

func (controller *DefaultController) Reset(context echo.Context) error {
	helpers.Database().Init();
	return context.String(http.StatusOK, "DB Reset Has Just Been Done!");
}