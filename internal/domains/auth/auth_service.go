package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mmcomp/go-rest-mysql-base/internal/domains/menu"
	"github.com/mmcomp/go-rest-mysql-base/internal/domains/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User = user.User

type AuthService struct {
	SecretKey        string
	DB               *gorm.DB
	GroupMenuService GroupMenuServiceInterface
	MenuService      MenuServiceInterface
}

func NewAuthService(secretKey string, db *gorm.DB, groupMenuService GroupMenuServiceInterface, menuService MenuServiceInterface) *AuthService {
	return &AuthService{SecretKey: secretKey, DB: db, GroupMenuService: groupMenuService, MenuService: menuService}
}

func (s *AuthService) GenerateAuthToken(userId uint, groupId uint) (LoginResponse, error) {
	authExp := time.Now().Add(time.Hour * 24 * 7).Unix()
	authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,
		"exp":   authExp,
		"type":  "auth",
		"group": groupId,
	})
	authTokenString, err := authToken.SignedString([]byte(s.SecretKey))
	if err != nil {
		fmt.Println("signing error", err)
		return LoginResponse{}, errors.New("signing error")
	}

	refExp := time.Now().Add(time.Hour * 24 * 256).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,
		"exp":   refExp,
		"type":  "refresh",
		"group": groupId,
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(s.SecretKey))
	if err != nil {
		fmt.Println("signing error", err)
		return LoginResponse{}, errors.New("signing error")
	}

	return LoginResponse{AuthToken: "Bearer " + authTokenString, AuthTokenExp: authExp, RefreshToken: "Bearer " + refreshTokenString, RefreshTokenExp: refExp}, nil
}

func (s *AuthService) GetUserByUsername(ctx context.Context, username string) (User, error) {
	var user User
	err := s.DB.WithContext(ctx).Preload("Group").Where("username = ?", username).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func HashPassword(password string) (string, error) {
	const cost = 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return strings.Replace(string(hash), "$2a$", "$2y$", 1), nil
}

func VerifyLaravelHash(password, laravelHash string) bool {
	if strings.HasPrefix(laravelHash, "$2y$") {
		laravelHash = "$2a$" + laravelHash[4:]
	}

	err := bcrypt.CompareHashAndPassword([]byte(laravelHash), []byte(password))
	return err == nil
}

func (s *AuthService) CheckUserPassword(ctx context.Context, username string, password string) (bool, User, error) {
	user, err := s.GetUserByUsername(ctx, username)
	if err != nil {
		return false, User{}, err
	}

	return VerifyLaravelHash(password, user.Password), user, nil
}

func (s *AuthService) GetUserMenus(ctx context.Context, user User) ([]menu.MenuNode, error) {
	groupMenus, err := s.GroupMenuService.GetAGroupMenus(ctx, user.GroupId)
	if err != nil {
		return nil, err
	}
	return s.MenuService.ArrangeMenusTreeLikeLike(groupMenus), nil
}

func (s *AuthService) Login(ctx context.Context, username string, password string) (UserLoginResponse, error) {
	valid, user, err := s.CheckUserPassword(ctx, username, password)
	if err != nil {
		return UserLoginResponse{}, err
	}
	if !valid {
		return UserLoginResponse{}, errors.New("invalid username or password")
	}
	loginResponse, err := s.GenerateAuthToken(user.ID, user.GroupId)
	if err != nil {
		return UserLoginResponse{}, err
	}
	menus, err := s.GetUserMenus(ctx, user)
	if err != nil {
		return UserLoginResponse{}, err
	}
	userDto := UserDto{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
	if user.Group != nil {
		userDto.Group = &GroupDto{ID: user.Group.ID, Name: user.Group.Name, Type: user.Group.Type}
	}
	return UserLoginResponse{LoginResponse: loginResponse, User: userDto, Menus: menus}, nil
}
