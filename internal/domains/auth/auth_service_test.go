package auth

import (
	"context"
	"testing"

	db_test "github.com/mmcomp/go-rest-mysql-base/internal/db_test"
	"github.com/mmcomp/go-rest-mysql-base/internal/domains/group"
	"github.com/mmcomp/go-rest-mysql-base/internal/domains/menu"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var database, cleanupDB = db_test.PrepareDatabase()
var ctx = context.Background()
var testuser = User{
	FirstName: "test",
	LastName:  "test",
	Username:  "test",
	Password:  "",
	GroupId:   1,
	Group:     &group.Group{ID: 1, Name: "test", Type: "test"},
}
var gms = &MockGroupMenuService{}
var mms = &MockMenuService{}

type MockGroupMenuService struct {
	mock.Mock
}

func (mgms *MockGroupMenuService) GetAGroupMenus(ctx context.Context, groupID uint) ([]menu.Menu, error) {
	mgms.Called(ctx, groupID)
	return []menu.Menu{}, nil
}

type MockMenuService struct {
	mock.Mock
}

func (mms *MockMenuService) ArrangeMenusTreeLikeLike(menus []menu.Menu) []menu.MenuNode {
	mms.Called(menus)
	return []menu.MenuNode{}
}

var authService = NewAuthService("secret", database, gms, mms)

func TestGetMD5Hash(t *testing.T) {
	hash, err := HashPassword("123456")
	require.NoError(t, err)
	require.True(t, VerifyLaravelHash("123456", hash))
}

func InitializeTestUser(t *testing.T) {
	cleanupDB()
	testuser.Password, _ = HashPassword("123456")
	database.Create(&testuser)
}

func TestGetUserByUsername(t *testing.T) {
	InitializeTestUser(t)
	user, err := authService.GetUserByUsername(ctx, "test")
	require.NoError(t, err)
	require.Equal(t, user.Username, "test")
}

func TestCheckUserPassword(t *testing.T) {
	InitializeTestUser(t)
	valid, user, err := authService.CheckUserPassword(ctx, "test", "123456")
	require.NoError(t, err)
	require.True(t, valid)
	require.Equal(t, user.Username, "test")
}

func TestGetUserMenus(t *testing.T) {
	InitializeTestUser(t)
	gms.Calls = []mock.Call{}
	mms.Calls = []mock.Call{}
	gms.On("GetAGroupMenus", ctx, testuser.GroupId).Return([]menu.Menu{}, nil)
	mms.On("ArrangeMenusTreeLikeLike", []menu.Menu{}).Return([]menu.MenuNode{}, nil)
	menus, err := authService.GetUserMenus(ctx, testuser)
	require.NoError(t, err)
	require.True(t, gms.AssertNumberOfCalls(t, "GetAGroupMenus", 1))
	require.True(t, mms.AssertNumberOfCalls(t, "ArrangeMenusTreeLikeLike", 1))
	require.Equal(t, menus, []menu.MenuNode{})
}

func TestLogin(t *testing.T) {
	InitializeTestUser(t)
	gms.Calls = []mock.Call{}
	mms.Calls = []mock.Call{}
	gms.On("GetAGroupMenus", ctx, testuser.GroupId).Return([]menu.Menu{}, nil)
	mms.On("ArrangeMenusTreeLikeLike", []menu.Menu{}).Return([]menu.MenuNode{}, nil)
	res, err := authService.Login(ctx, "test", "123456")
	require.NoError(t, err)
	require.Equal(t, res.User.FirstName, "test")
	require.Equal(t, res.User.LastName, "test")
	require.Equal(t, res.User.Group.ID, uint(1))
	require.Equal(t, res.User.Group.Name, "test")
	require.Equal(t, res.User.Group.Type, "test")
	require.Equal(t, res.Menus, []menu.MenuNode{})
	require.True(t, gms.AssertNumberOfCalls(t, "GetAGroupMenus", 1))
	require.True(t, mms.AssertNumberOfCalls(t, "ArrangeMenusTreeLikeLike", 1))
}
