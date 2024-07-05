package token

import (
	"api-gateway-service/generated/users"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	singingKey = "jwt-token-secret"
)

func GeneratedJWTToken(req *users.LoginResponse) error {
	token := *jwt.New(jwt.SigningMethodHS256)

	// payload
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = req.UserId
	claims["username"] = req.Username
	claims["email"] = req.Email
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()

	newToken, err := token.SignedString([]byte(singingKey))

	if err != nil {
		log.Println(err)
		return err
	}

	req.Token = newToken

	return nil
}

func ExtractClaims(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(singingKey), nil
	})

	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return err
	}

	return nil
}
