package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func getSecretKey() []byte {

	// Clave secreta para firmar el token
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	value, _ := os.LookupEnv("SECRET")
	return []byte(value)
}

var secretKey = getSecretKey()

// Estructura del JWT
type Claims struct {
	Userid         string `json:"userid"`
	Userpermission string `json:"userpermission"`
	jwt.RegisteredClaims
}

// Generar un nuevo token JWT
func GenerarToken(userid string) (string, error) {
	expirationTime := time.Now().Add(time.Hour)

	claims := &Claims{
		Userid:         userid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
