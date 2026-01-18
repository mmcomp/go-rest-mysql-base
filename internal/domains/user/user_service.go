package user

import (
	"context"

	"github.com/mmcomp/go-rest-mysql-base/common"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetUsers(ctx context.Context, request GetUsersRequest) (GetUsersResponse, error) {
	var users []User
	var total int64
	err := s.db.WithContext(ctx).Model(&User{}).Count(&total).Error
	if err != nil {
		return GetUsersResponse{}, err
	}
	err = s.db.WithContext(ctx).Offset((request.Page - 1) * request.PageSize).Limit(request.PageSize).Find(&users).Error
	if err != nil {
		return GetUsersResponse{}, err
	}
	usersDto := make([]UserDto, len(users))
	for i, user := range users {
		usersDto[i] = user.ToDto()
	}
	return GetUsersResponse{Data: usersDto, PaginatedResponseDTO: common.PaginatedResponseDTO{Page: request.Page, PageSize: request.PageSize, Total: int(total)}}, nil
}
