package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

/**
For those new to GO (Refer https://www.golang-book.com/books/intro/8)
	& and * is used in GO as pass by reference (Remember favorite chapter called Pointers in C?)
	& is generally mostly used in assignment, to create and return a pointer for a new value
	* is mostly used for declaration of, and access to the values that pointers point to (like those created in 1)
**/

type JWTCustomClaim struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

var jwtSecret = []byte(getJwtSecret())

func getJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "supersecret"
	}
	return secret
}

func JwtGenerate(ctx context.Context, userId string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTCustomClaim{
		ID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	})

	token, err := t.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return token, nil
}

func JwtValidate(ctx context.Context, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JWTCustomClaim{}, func (t *jwt.Token) (interface{}, error){
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}
		return jwtSecret, nil
	})
}
