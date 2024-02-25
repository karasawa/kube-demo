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
}

func InitRouting(e *echo.Echo, uc IUserController) {
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	pool := Pool{
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
	}
	
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, pool)
	})
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})
	e.GET("/user/:id", uc.GetUser)
}
