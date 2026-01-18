package auth

import "github.com/gin-gonic/gin"

func SetupAuthRoutes(authorized *gin.RouterGroup, unauthenticated *gin.RouterGroup, authController *AuthController) {
	authorized.GET("/api/v1/auth/ping", authController.Ping)
	authorized.GET("/api/v1/auth/refresh-token", authController.RefreshToken)
	unauthenticated.POST("/api/v1/auth/login", authController.Login)
}
