package jwtkeyfunc

import (
	"fmt"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v5"
)

func validateJWT(tokenString string, jwksURL string) (*jwt.Token, error) {
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{})
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, jwks.Keyfunc)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}
