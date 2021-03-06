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
	r.HandleFunc("/login/{adminid}/createticket", controllers.CreateTicket).Methods("POST")
	r.HandleFunc("/login/{adminid}/updateticket/{ticketid}", controllers.UpdateTicket).Methods("PUT")
	r.HandleFunc("/login/{adminid}/gettickets/{timings}", controllers.GetTicket).Methods("GET")
	r.HandleFunc("/login/{adminid}/deletetickets/{ticketid}", controllers.DeleteTicket).Methods("DELETE")
	r.HandleFunc("/login/{adminid}/userdetailstickets/{ticketid}", controllers.UserDetailsTicket).Methods("GET")
	r.HandleFunc("/login/{adminid}/markticketexpired/{ticketid}", controllers.MarkTicketExpired).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}
