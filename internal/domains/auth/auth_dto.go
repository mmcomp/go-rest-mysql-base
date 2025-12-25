package auth

import (
	"github.com/mmcomp/go-rest-mysql-base/internal/domains/menu"
)

type GroupDto struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type UserDto struct {
	ID        uint      `json:"id"`
	Group     *GroupDto `json:"group"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AuthToken       string `json:"auth_token"`
	AuthTokenExp    int64  `json:"auth_token_exp"`
	RefreshToken    string `json:"refresh_token"`
	RefreshTokenExp int64  `json:"refresh_token_exp"`
}

type UserLoginResponse struct {
	LoginResponse
	User  UserDto         `json:"user"`
	Menus []menu.MenuNode `json:"menus"`
}
