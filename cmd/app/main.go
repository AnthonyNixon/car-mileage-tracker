package main

import (
	"AnthonyNixon/car-mileage-tracker/cmd/handlers/events"
	"AnthonyNixon/car-mileage-tracker/cmd/handlers/up"
	"AnthonyNixon/car-mileage-tracker/cmd/services/router"
	"flag"
	"fmt"
	"log"
	"os"
)

var PORT = ""

func init() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
}

func main() {
	releaseModePtr := flag.Bool("release", false, "release mode enabled. If false, debug mode is active.")
	flag.Parse()

	router := router.New(*releaseModePtr)

	up.AddUpV1(router)
	events.AddEventsV1(router)

	log.Printf("Running car-mileage-tracker API on :%s...", PORT)
	err := router.Run(fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatal(err.Error())
	}
}
