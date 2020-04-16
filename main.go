package main

import (
	"log"
	"net/http"

	"github.com/dshoulders/goapi/controllers"
	"github.com/dshoulders/goapi/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	r.Use(middleware.JwtAuthentication) //attach JWT auth middleware

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)

	api.HandleFunc("/list", controllers.GetLists).Methods(http.MethodGet)
	api.HandleFunc("/list", controllers.CreateList).Methods(http.MethodPost)
	api.HandleFunc("/list/{listId}", controllers.GetList).Methods(http.MethodGet)
	api.HandleFunc("/list/{listId}/listitem", controllers.GetListItems).Methods(http.MethodGet)
	api.HandleFunc("/list/{listId}/listitem", controllers.CreateListItem).Methods(http.MethodPost)
	api.HandleFunc("/listitem/{itemId}", controllers.GetListItem).Methods(http.MethodGet)
	api.HandleFunc("/listitem/{itemId}", controllers.UpdateListItem).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":8000", r))
}
