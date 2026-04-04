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

func CreateJWT(jwtSecretKey []byte, data Payload) string {
	claims := jwt.MapClaims{
		"sub":        data.Sub,
		"first_name": data.FirstName,
		"role":       data.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return ""
	}
	return tokenString
}

func VerifyJWT(jwtSecretKey []byte, tokenStr string) (Payload, bool) {
	token, err := jwt.Parse(
		tokenStr,
		func(t *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		},
	)
	if err != nil {
		return Payload{}, false
	}
	if token.Valid == false {
		return Payload{}, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Payload{}, false
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return Payload{}, false
	}
	firstName, ok := claims["first_name"].(string)
	if !ok {
		return Payload{}, false
	}
	role, ok := claims["role"].(string)
	if !ok {
		return Payload{}, false
	}

	data := Payload{
		Sub:       sub,
		FirstName: firstName,
		Role:      role,
	}

	return data, true
}
