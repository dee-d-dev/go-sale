package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var atSecretKey = []byte(os.Getenv("ACCESS_SECRET"))
var rtSecretKey = []byte(os.Getenv("REFRESH_SECRET"))

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func GenerateToken(email string) (Tokens, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	})

	accessTokenString, err := accessToken.SignedString(atSecretKey)

	if err != nil {
		return Tokens{}, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	})

	refreshTokenString, err := refreshToken.SignedString(rtSecretKey)

	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
