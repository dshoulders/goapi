package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dshoulders/goapi/models"
	"github.com/dshoulders/goapi/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtAuthentication handles all requests checking for a wellformed and valid JWT
// Some endpoints such as login are allowed to pass through without checking for a JWT
var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/login"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path         //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			response := models.CreateErrorResponse("Missing auth token")
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response, http.StatusUnauthorized)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response := models.CreateErrorResponse("Invalid/Malformed auth token")
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response, http.StatusUnauthorized)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("AUTH_TOKEN")), nil
		})

		if tk.ExpiresAt < time.Now().Unix() { //Token has expired
			response := models.CreateErrorResponse("Token has expired")
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response, http.StatusUnauthorized)
			return
		}

		if err != nil { //Malformed token, returns with http code 403 as usual
			response := models.CreateErrorResponse("Malformed authentication token")
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response, http.StatusUnauthorized)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response := models.CreateErrorResponse("Token is not valid")
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response, http.StatusUnauthorized)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Sprintf("User %", tk.Username) //Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
