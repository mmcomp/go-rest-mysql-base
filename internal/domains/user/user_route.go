package user

import "github.com/gin-gonic/gin"

func SetupUserRoutes(router *gin.RouterGroup, userController *UserController) {
	router.GET("/api/v1/users", userController.GetUsers)
}
