package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte(GetSecretKeyJWT())

func GenerateJWT(u string) (string, error) {
	// Crea un nuevo token JWT con el método de firma adecuado
	token := jwt.New(jwt.SigningMethodHS256)

	// Configura los claims del token
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Hour).Unix()
	claims["authorized"] = true
	claims["user"] = u

	// Firma el token con la clave secreta
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyTokenJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("Token") // Utiliza Get en lugar de un índice 0
		if token != "" {
			parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Invalid signature method: %v", token.Header["alg"])
				}
				return secretKey, nil
			})

			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err2 := writer.Write([]byte("Unauthorized due to error parsing the JWT"))
				if err2 != nil {
					return
				}
			} else if parsedToken.Valid {
				endpointHandler(writer, request)
			} else {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err := writer.Write([]byte("Unauthorized due to invalid token"))
				if err != nil {
					return
				}
			}
		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			_, err := writer.Write([]byte("Unauthorized: Missing Token"))
			if err != nil {
				return
			}
		}
	})
}
