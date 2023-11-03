package main

import (
	"JurrassicParkAPI/app/bootstrap"
	"JurrassicParkAPI/app/controller"
	"JurrassicParkAPI/app/pkg"
	"JurrassicParkAPI/app/repository"
	"JurrassicParkAPI/app/router"
	"JurrassicParkAPI/app/service"
	"JurrassicParkAPI/config"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/fx"
)

func main() {
	config.InitLog()
	app := fx.New(
		config.Module,
		pkg.Module,
		controller.Module,
		router.Module,
		service.Module,
		repository.Module,
		bootstrap.Module, //Gin Startup is invoked in the bootstrap module.
	)

	app.Run()
}
