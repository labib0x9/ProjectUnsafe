package utils

import (
	"github.com/golang-jwt/jwt"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub       string `json:"sub"`
	FirstName string `json:"first_name"`
	Role      string `json:"role"`
}

func CreateJWT(jwtSecretKey string, data Payload) string {
	claims := jwt.MapClaims{
		"sub":        data.Sub,
		"first_name": data.FirstName,
		"role":       data.Role,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenStr, err := token.SignedString(jwtSecretKey)
	if err != nil {
		//
		return ""
	}
	return tokenStr
}
