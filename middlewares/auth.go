package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetToken(headers http.Header) (string, error) {
	for hkey, values := range headers {
		key := strings.ToLower(hkey)
		if key == "authorization" && len(values) == 1 {
			token := strings.TrimSpace(values[0])
			tokenParts := strings.Split(token, " ")
			if len(tokenParts) > 2 || len(tokenParts) < 1 {
				return "", fmt.Errorf("invalid token format: %s", token)
			}
			if len(tokenParts) == 2 {
				token = tokenParts[1]
			}
			return token, nil
		}
	}

	return "", errors.New("authentication token not found")
}

func ParseToken(tokenString, secretKey string) (uint, string, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, "", err
	}
	if !token.Valid {
		return 0, "", fmt.Errorf("token is not valid")
	}

	// Extract user_id from claims
	userIdFloat, ok := claims["sub"].(float64)
	if !ok {
		return 0, "", fmt.Errorf("userId claim is not found or not a number")
	}
	groupName, ok := claims["group"].(string)
	if !ok {
		// return 0, "", fmt.Errorf("group claim is not found or not a string")
		groupName = ""
	}
	return uint(userIdFloat), groupName, nil
}

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := GetToken(c.Request.Header)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "authentication failed, token not found"})
			c.Abort()
			return
		}
		token = strings.ReplaceAll(token, "Bearer ", "")
		userId, groupName, err := ParseToken(token, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		c.Set("userId", userId)
		c.Set("group", groupName)

		c.Next()
	}
}
