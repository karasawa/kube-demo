package main

import (
	"kube/controller"
	"kube/repository"
	"kube/usercase"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	e := echo.New()
	ur := repository.NewUserRepository()
	uu := usercase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)
	controller.InitRouting(e, uc)
	e.Logger.Fatal(e.Start(":8080"))
}
