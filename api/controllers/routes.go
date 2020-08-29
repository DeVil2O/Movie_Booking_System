package controllers

import (
	"github.com/DeVil2O/moviebookingsystem/api/middlewares"
)

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/login", s.Login).Methods("POST")

	s.Router.HandleFunc("/users", middlewares.SetMiddlewareAuthentication(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/user", middlewares.SetMiddlewareAuthentication(s.CRUDUser)).Methods("POST")
	s.Router.HandleFunc("/user", middlewares.SetMiddlewareAuthentication(s.CRUDUser)).Methods("PUT")
	s.Router.HandleFunc("/user", middlewares.SetMiddlewareAuthentication(s.CRUDUser)).Methods("DELETE")

	s.Router.Use(middlewares.SetMiddlewareJSON)
}
