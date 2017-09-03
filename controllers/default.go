package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

type DefaultController struct {

}

type JsonResponse struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	Data    interface{} `json:"data"`
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

func (controller *DefaultController) Index(context echo.Context) error {
	return context.JSON(http.StatusOK, JsonResponse{Code: 0, Message: "Hello World! It's working..."})
}