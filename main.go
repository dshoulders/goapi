package main

import (
	"log"
	"net/http"

	"github.com/dshoulders/goapi/controllers"
	"github.com/dshoulders/goapi/middleware"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.JwtAuthentication) //attach JWT auth middleware

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)

	api.HandleFunc("/list", controllers.GetLists).Methods(http.MethodGet)
	api.HandleFunc("/list", controllers.CreateList).Methods(http.MethodPost)
	api.HandleFunc("/list/{listId}", controllers.GetList).Methods(http.MethodGet)
	api.HandleFunc("/list/{listId}/listitem", controllers.ListItems).Methods(http.MethodGet)
	api.HandleFunc("/listitem/{itemId}", controllers.ListItem).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", r))
}
