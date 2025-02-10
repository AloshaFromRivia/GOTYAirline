package main

import (
	"UltimateDeluxeSuperAirlineManagerGOTY"
	"UltimateDeluxeSuperAirlineManagerGOTY/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(airlineManager.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("%v", err)
	}
}
