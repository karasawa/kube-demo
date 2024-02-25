package controller

import (
	"kube/usercase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	GetUser(echo.Context) error
}

type userController struct {
	uu usercase.IUserUsecase
}

func NewUserController(uu usercase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) GetUser(c echo.Context) error {
	id := c.Param("id")
	res, err := uc.uu.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, *res)
}
