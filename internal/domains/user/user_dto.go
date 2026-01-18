package user

import "github.com/mmcomp/go-rest-mysql-base/common"

type UserDto struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GetUsersRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type GetUsersResponse struct {
	common.PaginatedResponseDTO
	Data []UserDto `json:"data"`
}
