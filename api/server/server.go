package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DeVil2O/moviebookingsystem/api/controllers"
	"github.com/gorilla/mux"
)

func Run() {
	fmt.Println("Chirag Garg")
	r := mux.NewRouter()
	r.HandleFunc("/register", controllers.RegisterHandler).
		Methods("POST")
	r.HandleFunc("/login", controllers.LoginHandler).
		Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
