package main

import (
	"UltimateDeluxeSuperAirlineManagerGOTY"
	"UltimateDeluxeSuperAirlineManagerGOTY/pkg/handler"
	"UltimateDeluxeSuperAirlineManagerGOTY/repository"
	"UltimateDeluxeSuperAirlineManagerGOTY/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(airlineManager.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("%v", err)
	}
}
