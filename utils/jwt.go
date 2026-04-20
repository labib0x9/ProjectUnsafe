package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labib0x9/ProjectUnsafe/model"
)

type Payload struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func CreateJWT(jwtSecretKey []byte, data model.User) (string, error) {
	claims := Payload{
		Fullname: data.Fullname,
		Email:    data.Email,
		Role:     data.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   data.Id.String(),
			Issuer:    "projectpdf",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func VerifyJWT(jwtSecretKey []byte, tokenStr string) (Payload, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Payload{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("wrong signing method: %v", t.Header["alg"])
			}
			return jwtSecretKey, nil
		},
	)

	if err != nil {
		return Payload{}, err
	}

	claims, ok := token.Claims.(*Payload)
	if !ok {
		return Payload{}, errors.New("invalid claims")
	}

	return *claims, nil
}
