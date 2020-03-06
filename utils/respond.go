package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func Respond(w http.ResponseWriter, responseData interface{}, status int) {
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
