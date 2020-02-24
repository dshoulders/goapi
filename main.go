package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

// Message Shape
type Message struct {
	Message string `json:"message"`
}

func respond(w http.ResponseWriter, responseData interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(responseData)
	if err == nil {
		w.WriteHeader(status)
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(
			[]byte(
				fmt.Sprintf(
					`{"message": "Type: %s cannot be serialised"}`,
					reflect.TypeOf(responseData),
				),
			),
		)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Message: "get called",
	}
	respond(w, message, http.StatusOK)
}

func post(w http.ResponseWriter, r *http.Request) {
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Println(err)
	}

	respond(w, message, http.StatusCreated)
}

func params(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	userID := -1
	var err error
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a user Id"}`))
			return
		}
	}

	commentID := -1
	if val, ok := pathParams["commentID"]; ok {
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a comment Id"}`))
			return
		}
	}

	query := r.URL.Query()
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s"}`, userID, commentID, location)))
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)

	api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", r))
}
