package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dshoulders/goapi/models"
	"github.com/dshoulders/goapi/services"
	"github.com/dshoulders/goapi/utils"
)

type LoginCredentials struct {
	Username string
	Password string
}

type LoginResponse struct {
	Success bool `json:"success"`
	services.TokenPair
}

// Login - Handles login requests and returns a JWT if succesful
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials LoginCredentials
	var status int
	var tokenPair services.TokenPair

	_ = json.NewDecoder(r.Body).Decode(&credentials)
	tokenPair, err := services.Authenticate(credentials.Username, credentials.Password)

	if err != nil {
		response := models.CreateErrorResponse(err.Error())
		status = http.StatusUnauthorized
		utils.Respond(w, response, status)
	} else {
		response := models.CreateSuccessResponse(tokenPair)
		status = http.StatusOK
		utils.Respond(w, response, status)
	}
}
