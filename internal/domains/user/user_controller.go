package user

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetUsers(ctx context.Context, request GetUsersRequest) (GetUsersResponse, error)
}

type UserController struct {
	userService UserServiceInterface
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{userService: NewUserService(db)}
}

// Get Users godoc
//
//	@Summary	Get Users
//	@Security	AuthToken
//	@Schemes
//	@Description	Get Users
//	@Tags			User
//	@Param			page		query		int	false	"Page"		default(1)
//	@Param			page_size	query		int	false	"Page Size"	default(10)
//	@Success		200			{object}	GetUsersResponse
//	@Router			/users [get]
func (c *UserController) GetUsers(ctx *gin.Context) {
	var request GetUsersRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	if request.Page < 1 {
		request.Page = 1
	}
	if request.PageSize < 1 {
		request.PageSize = 10
	}
	users, err := c.userService.GetUsers(ctx.Request.Context(), request)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get users"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
