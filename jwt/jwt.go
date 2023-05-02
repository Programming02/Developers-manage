package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtgo "github.com/dgrijalva/jwt-go"
	"strings"
)

func ExtractClaims(t string, signKey []byte) (jwt.MapClaims, error) {
	claims := jwtgo.MapClaims{}
	if t == "" {
		claims["role"] = "unauthorized"
		return claims, nil
	}
	if strings.Contains(t, "Basic") {
		claims["role"] = "unauthorized"
		return claims, nil
	}
	token, err := jwtgo.ParseWithClaims(t, claims, func(token *jwtgo.Token) (interface{}, error) {
		return signKey, nil
	})

	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !(ok && token.Valid) {
		err = fmt.Errorf("JWT token bad")
		return nil, err
	}
	return claims, nil
}
