package utils

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetRequestParamAsInt(name string, r *http.Request) (int, error) {
	pathParams := mux.Vars(r)

	paramValue := -1
	var err error
	if val, ok := pathParams[name]; ok {
		paramValue, err = strconv.Atoi(val)
	}
	return paramValue, err
}
