package thirdparties

import (
	"os"
)

var jwtUtil *JWTUtil

func InitThirdParties() {
	jwtUtil = NewJWTUtil(os.Getenv("AUTH_SECRET_KEY"))
}

func GetJWTUtil() *JWTUtil {
	if jwtUtil == nil {
		InitThirdParties()
	}
	return jwtUtil
}
