package auth

import "github.com/gin-gonic/gin"

func SetupAuthRoutes(authorized *gin.RouterGroup, unauthenticated *gin.RouterGroup, authController *AuthController) {
	authorized.GET("/api/v1/admin/ping", authController.Ping)
	authorized.GET("/api/v1/admin/auth/refresh-token", authController.RefreshToken)
	unauthenticated.POST("/api/v1/admin/auth/login", authController.Login)
}
