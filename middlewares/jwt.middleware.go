package middlewares

import (
	"fmt"
	"log"
	"net/http"
	// "strings"

	"os"

	"github.com/dgrijalva/jwt-go"
)

func TokenMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractToken(r)

		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unauthorized")

			}

			return []byte(os.Getenv("AT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			log.Printf("Error while parsing token: %v", err)
			http.Error(w, "Unauthorised: Invalid token", http.StatusUnauthorized)

			return
		}

		next.ServeHTTP(w, r)

	})
}

func extractToken(r *http.Request) string {
	// token := r.Header.Get("Authorization")
	token, err := r.Cookie("access_token")

	if err != nil {
		log.Fatal(err)
	}

	if token.Value == "" {
		return ""
	}

	// parts := strings.Split(token, " ")
	// if len(parts) != 2 || parts[0] != "Bearer" {
	// 	return ""
	// }

	// Return the token without the "Bearer " prefix
	// return parts[1]
	return token.Value

}
