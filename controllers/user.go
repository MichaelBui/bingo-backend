package controllers

import (
	"github.com/labstack/echo"
	"github.com/michaelbui/bingo-backend/entities"
	"net/http"
	"github.com/michaelbui/bingo-backend/models"
)

type UserController struct {}

var (
	user *UserController
)

func User() *UserController {
	if user == nil {
		user = &UserController{}
	}
	return user
}

func (u *UserController) Add(context echo.Context) error {
	user := &entities.UserEntity{}
	postData := struct{
		*entities.UserEntity
		Password string `json:"password"`
		PasswordConfirm string `json:"passwordConfirm"`
	}{user, "", ""}

	err := context.Bind(&postData)
	if err != nil {
		return context.JSON(http.StatusBadRequest, JsonResponse{Code: 1, Message: err.Error()})
	}

	if !user.VerifyPassword(postData.Password, postData.PasswordConfirm) {
		return context.JSON(http.StatusBadRequest, JsonResponse{Code: 2, Message: "Unmatched Passwords"})
	}

	user.SetPassword(postData.Password)
	*user, err = models.User().Add(*user)
	if err != nil {
		return context.JSON(http.StatusBadRequest, JsonResponse{Code: 2, Message: err.Error()})
	}

	return context.JSON(http.StatusOK, JsonResponse{Data: user})
}

func (u *UserController) Get(context echo.Context) error {
	email := context.Param("email")
	password := context.Request().Header.Get("X-GUP")
	e, err := models.User().Get(email, password)
	if err != nil {
		return context.JSON(http.StatusUnauthorized, JsonResponse{Code: 1, Message: err.Error(), Data: e})
	}
	return context.JSON(http.StatusOK, JsonResponse{Data: e})
}

func (u *UserController) UpdateNumbers(context echo.Context) error {
	email := context.Param("email")
	password := context.Request().Header.Get("X-GUP")
	e, err := models.User().Get(email, password)
	if err != nil {
		return context.JSON(http.StatusUnauthorized, JsonResponse{Code: 1, Message: err.Error()})
	}

	numbers := []int{}
	err = context.Bind(&numbers)
	if err != nil {
		return context.JSON(http.StatusBadRequest, JsonResponse{Code: 2, Message: err.Error()})
		return context.JSON(http.StatusBadRequest, JsonResponse{Code: 2, Message: err.Error()})
	}

	err = models.User().UpdateNumbers(&e, numbers)
	if err != nil {
		return context.JSON(http.StatusBadRequest, JsonResponse{Code: 3, Message: err.Error()})
	}
	return context.JSON(http.StatusOK, JsonResponse{Data: e})
}