package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	m "github.com/dshoulders/goapi/models"
	u "github.com/dshoulders/goapi/utils"

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
			response := m.Message{"Missing auth token"}
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response, http.StatusForbidden)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response := m.Message{"Invalid/Malformed auth token"}
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response, http.StatusForbidden)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &m.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("AUTH_TOKEN")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			response := m.Message{"Malformed authentication token"}
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response, http.StatusForbidden)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response := m.Message{"Token is not valid."}
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response, http.StatusForbidden)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Sprintf("User %", tk.Username) //Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}
