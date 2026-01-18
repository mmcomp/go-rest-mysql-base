package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserDomain(authorized *gin.RouterGroup, db *gorm.DB) {
	userController := NewUserController(db)
	SetupUserRoutes(authorized, userController)
}
