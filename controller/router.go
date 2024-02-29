package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRouting(e *echo.Echo, uc IUserController) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/user/:id", uc.GetUserById)
	e.GET("/users", uc.GetUserList)
	e.POST("/user", uc.CreateUser)
	e.PUT("/user/:id", uc.UpdateUser)
}
