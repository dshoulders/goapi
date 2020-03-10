package services

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/dshoulders/goapi/models"
)

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func GenerateToken(tokenClaims models.Token) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenClaims)
	return token.SignedString([]byte(os.Getenv("AUTH_TOKEN")))
}

// Authenticate - Authenticate a user based on username and password
func Authenticate(username string, password string) (TokenPair, error) {

	user, err := GetUser(username)

	if err != nil {
		return TokenPair{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return TokenPair{}, errors.New("Invalid login credentials. Please try again")
	}

	if err != nil {
		return TokenPair{}, err
	}

	now := time.Now()
	after15Minutes := now.Add(time.Minute * 15)
	after30Days := now.Add(time.Hour * 24 * 30)

	//Create access JWT token
	accessTokenClaims := models.Token{
		UserId: user.Id,
	}
	accessTokenClaims.ExpiresAt = after15Minutes.Unix()
	accessTokenString, _ := GenerateToken(accessTokenClaims)

	//Create refresh JWT token
	refreshTokenClaims := models.Token{
		UserId: user.Id,
	}
	refreshTokenClaims.ExpiresAt = after30Days.Unix()
	refreshTokenString, _ := GenerateToken(refreshTokenClaims)

	tokenPair := TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	return tokenPair, nil
}
