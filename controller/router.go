package controller

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type Pool struct {
	DB_USER string
	DB_PASSWORD string
	DB_NAME string
	INSTANCE_UNIX_SOCKET string
}

func InitRouting(e *echo.Echo, uc IUserController) {
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	INSTANCE_UNIX_SOCKET := os.Getenv("INSTANCE_UNIX_SOCKET")
	pool := Pool{
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		INSTANCE_UNIX_SOCKET,
	}
	
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, pool)
	})
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})
	e.GET("/user/:id", uc.GetUserById)
	e.GET("/users", uc.GetUserList)
	e.POST("/user", uc.CreateUser)
	e.PUT("/user/:id", uc.UpdateUser)
}
