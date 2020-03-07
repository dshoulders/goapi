package models

import "github.com/dgrijalva/jwt-go"

/*
JWT claims struct
*/
type Token struct {
	UserId   int32 `json:"userId"`
	Username int32 `json:"username"`
	jwt.StandardClaims
}
