package thirdparties

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTUtil struct {
	secretKey string
}

func NewJWTUtil(secretKey string) *JWTUtil {
	return &JWTUtil{secretKey: secretKey}
}

func (j *JWTUtil) ParseToken(tokenString string) (uint, string, time.Time, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return 0, "", time.Time{}, err
	}
	if !token.Valid {
		return 0, "", time.Time{}, fmt.Errorf("token is not valid")
	}

	// Extract user_id from claims
	userIdFloat, ok := claims["sub"].(float64)
	if !ok {
		return 0, "", time.Time{}, fmt.Errorf("userId claim is not found or not a number")
	}
	groupName, ok := claims["group"].(string)
	if !ok {
		return 0, "", time.Time{}, fmt.Errorf("group claim is not found or not a string")
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		return 0, "", time.Time{}, fmt.Errorf("exp claim is not found or not a number")
	}
	return uint(userIdFloat), groupName, time.Unix(int64(exp), 0), nil
}
