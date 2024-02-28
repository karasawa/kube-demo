package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Pool struct {
	DB_USER string
	DB_PASSWORD string
	DB_NAME string
	INSTANCE_UNIX_SOCKET string
}

func InitRouting(e *echo.Echo, uc IUserController) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/user/:id", uc.GetUserById)
	e.GET("/users", uc.GetUserList)
	e.POST("/user", uc.CreateUser)
	e.PUT("/user/:id", uc.UpdateUser)
}
