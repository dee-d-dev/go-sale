package utils

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// var atSecretKey string


type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func GenerateToken(email string) (Tokens, error) {
	AT_SECRET, exists := os.LookupEnv("AT_SECRET")

	if !exists {
		log.Fatal("AT_SECRET not set")
	}

	atSecretKeyExists := []byte(AT_SECRET)

	RT_SECRET, exists := os.LookupEnv("RT_SECRET")

	if !exists {
		log.Fatal("RT_SECRET not set")
	}
	rtSecretKeyExists := []byte(RT_SECRET)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	})

	accessTokenString, err := accessToken.SignedString(atSecretKeyExists)

	if err != nil {
		return Tokens{}, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	})

	refreshTokenString, err := refreshToken.SignedString(rtSecretKeyExists)

	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
