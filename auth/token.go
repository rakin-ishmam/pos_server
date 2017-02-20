package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rakin-ishmam/pos_server/apperr"
)

// New returns jwt token
func New(info Info, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserName": info.UserName,
		"Exp":      info.Exp.Format(time.RFC3339),
	})

	return token.SignedString([]byte(secret))
}

// Decode decodes jwt token
func Decode(tokenStr, secret string) (*Info, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		info := Info{UserName: claims["UserName"].(string)}
		info.Exp, err = time.Parse(time.RFC3339, claims["Exp"].(string))

		if err != nil {
			return nil, err
		}

		return &info, nil
	}

	return nil, apperr.Authentication{Where: "Token", Cause: apperr.StrInvalid}
}
