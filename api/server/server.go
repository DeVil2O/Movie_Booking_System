package server

import (
	"fmt"
	"log"

	"github.com/DeVil2O/moviebookingsystem/api/controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting environment, not coming through %v", err)
	} else {
		fmt.Println("We are getting the environment values")
	}

	server.Initialize()

	server.Run(":8080")

}
