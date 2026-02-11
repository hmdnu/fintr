package token

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hmdnu/fintr/pkg/env"
)

func GenerateToken(userId int) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   strconv.Itoa(userId),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(env.JWT_SECRET))
}

func VerifiyToken(token string) (*jwt.Token, bool) {
	claims := &jwt.RegisteredClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		return []byte(env.JWT_SECRET), nil
	})
	if err != nil || !jwtToken.Valid {
		return nil, false
	}
	return jwtToken, true
}
