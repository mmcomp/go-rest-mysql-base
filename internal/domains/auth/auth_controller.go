package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mmcomp/go-rest-mysql-base/internal/domains/menu"
	"gorm.io/gorm"
)

type AuthServiceInterface interface {
	GenerateAuthToken(userId uint, groupId uint) (LoginResponse, error)
	CheckUserPassword(ctx context.Context, username string, password string) (bool, User, error)
	Login(ctx context.Context, username string, password string) (UserLoginResponse, error)
}
type GroupMenuServiceInterface interface {
	GetAGroupMenus(ctx context.Context, groupID uint) ([]menu.Menu, error)
}
type MenuServiceInterface interface {
	ArrangeMenusTreeLikeLike(menus []menu.Menu) []menu.MenuNode
}
type AuthController struct {
	authService AuthServiceInterface
}

func NewAuthController(secretKey string, db *gorm.DB, groupMenuService GroupMenuServiceInterface, menuService MenuServiceInterface) *AuthController {
	return &AuthController{authService: NewAuthService(secretKey, db, groupMenuService, menuService)}
}

//	@BasePath	/api/v1

// Auth Ping godoc
//
//	@Summary	Ping
//	@Security	AuthToken
//	@Schemes
//	@Description	Ping
//	@Tags			Auth
//	@Success		200	{string}	string	"pong"
//	@Router			/auth/ping [get]
func (s *AuthController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

// Auth Refresh Token godoc
//
//	@Summary	Refresh Token
//	@Security	RefreshToken
//	@Schemes
//	@Description	Refresh Token
//	@Tags			Auth
//	@Success		200	{object}	UserLoginResponse
//	@Router			/auth/refresh-token [get]
func (s *AuthController) RefreshToken(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized, invalid userId"})
		return
	}

	groupId, exists := c.Get("group")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized, invalid groupName"})
		return
	}

	loginResponse, err := s.authService.GenerateAuthToken(userId.(uint), groupId.(uint))
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to generate auth token"})
		return
	}

	c.JSON(200, loginResponse)
}

// Auth Login godoc
//
//	@Summary	Login
//	@Schemes
//	@Description	Login
//	@Tags			Auth
//	@Param			request	body		LoginRequest	true	"Login request"
//	@Success		200		{object}	UserLoginResponse
//	@Router			/auth/login [post]
func (s *AuthController) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid username or password"})
		return
	}

	userLoginResponse, err := s.authService.Login(c.Request.Context(), request.Username, request.Password)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to login"})
		return
	}

	c.JSON(200, userLoginResponse)
}
