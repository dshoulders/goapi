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
	r.Use(middleware.JwtAuthentication) // attach JWT auth middleware

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)

	// Get all lists for user
	api.HandleFunc("/list", controllers.GetLists).Methods(http.MethodGet)
	// Create new list
	api.HandleFunc("/list", controllers.CreateList).Methods(http.MethodPost)
	// Get list details
	api.HandleFunc("/list/{listId}", controllers.GetList).Methods(http.MethodGet)
	// Delete list and all list items within list
	api.HandleFunc("/list/{listId}", controllers.DeleteList).Methods(http.MethodDelete)
	// Get all list items in list
	api.HandleFunc("/list/{listId}/listitem", controllers.GetListItems).Methods(http.MethodGet)
	// Create new list item
	api.HandleFunc("/list/{listId}/listitem", controllers.CreateListItem).Methods(http.MethodPost)
	// Get list item
	api.HandleFunc("/listitem/{itemId}", controllers.GetListItem).Methods(http.MethodGet)
	// Update list item
	api.HandleFunc("/listitem/{itemId}", controllers.UpdateListItem).Methods(http.MethodPut)
	// Delete list item
	api.HandleFunc("/listitem/{itemId}", controllers.DeleteListItem).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8000", r))
}
