package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dshoulders/goapi/controllers"
	"github.com/dshoulders/goapi/middleware"
	"github.com/dshoulders/goapi/models"
	"github.com/dshoulders/goapi/utils"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func get(w http.ResponseWriter, r *http.Request) {
	message := models.Message{
		Message: "get called",
	}
	utils.Respond(w, message, http.StatusOK)
}

func post(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Println(err)
	}

	utils.Respond(w, message, http.StatusCreated)
}

func params(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	userId := -1
	var err error
	if val, ok := pathParams["userId"]; ok {
		userId, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a user Id"}`))
			return
		}
	}

	commentId := -1
	if val, ok := pathParams["commentId"]; ok {
		commentId, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a comment Id"}`))
			return
		}
	}

	query := r.URL.Query()
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"userId": %d, "commentId": %d, "location": "%s"}`, userId, commentId, location)))
}

func main() {
	r := mux.NewRouter()
	r.Use(middleware.JwtAuthentication) //attach JWT auth middleware

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)

	api.HandleFunc("/listitem/{itemId}", controllers.ListItem).Methods(http.MethodGet)

	api.HandleFunc("/user/{userId}/comment/{commentId}", params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", r))
}
