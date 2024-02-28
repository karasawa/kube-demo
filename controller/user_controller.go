package controller

import (
	"kube/entity"
	"kube/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	GetUserById(echo.Context) error
	GetUserList(echo.Context) error
	CreateUser(echo.Context) error
	UpdateUser(echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) GetUserById(c echo.Context) error {
	id := c.Param("id")
	res, err := uc.uu.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, *res)
}

func (uc *userController) GetUserList(c echo.Context) error {
	res, err := uc.uu.GetUserList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (uc *userController) CreateUser(c echo.Context) error {
	user := entity.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	res, err := uc.uu.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, *res)
}

func (uc *userController) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := entity.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user.Id = id
	res, err := uc.uu.UpdateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, *res)
}
