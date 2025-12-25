package auth

import (
	"github.com/gin-gonic/gin"
	groupmenu "github.com/mmcomp/go-rest-mysql-base/internal/domains/group_menu"
	"github.com/mmcomp/go-rest-mysql-base/internal/domains/menu"
	"gorm.io/gorm"
)

func InitAuthDomain(authorized *gin.RouterGroup, unauthenticated *gin.RouterGroup, secretKey string, db *gorm.DB) {
	groupMenuService := groupmenu.NewGroupMenuService(db)
	menuService := menu.NewMenuService(db)
	authController := NewAuthController(secretKey, db, groupMenuService, menuService)
	SetupAuthRoutes(authorized, unauthenticated, authController)
}
