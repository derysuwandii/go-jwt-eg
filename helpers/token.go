package helpers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-jwt-eg/entities"
	"time"
)

var mySigningKey = []byte("kdfhlitgner89hg89hg978h34gb348g9b3g834gn9")

type MyCustomClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user *entities.User) (string, error) {
	claims := MyCustomClaims{
		user.ID, user.Name, user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ValidateToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	}

	return claims, nil
}
