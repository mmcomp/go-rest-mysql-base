package groupmenu

import (
	"context"
	"slices"

	"github.com/mmcomp/go-rest-mysql-base/internal/domains/menu"
	"gorm.io/gorm"
)

type GroupMenuService struct {
	db *gorm.DB
}

func NewGroupMenuService(db *gorm.DB) *GroupMenuService {
	return &GroupMenuService{db: db}
}

func (s *GroupMenuService) LoadGroupMenusIntoMemory(ctx context.Context) (map[uint][]menu.Menu, error) {
	groupMenusMap := make(map[uint][]menu.Menu)
	var groupMenus []GroupMenu
	err := s.db.WithContext(ctx).Preload("Menu").Find(&groupMenus).Error
	if err != nil {
		return nil, err
	}
	for _, groupMenu := range groupMenus {
		groupMenusMap[groupMenu.GroupID] = append(groupMenusMap[groupMenu.GroupID], *groupMenu.Menu)
	}
	return groupMenusMap, nil
}

func (s *GroupMenuService) GetGroupMenus(ctx context.Context, groupID uint) ([]GroupMenu, error) {
	var groupMenus []GroupMenu
	err := s.db.WithContext(ctx).Preload("Menu").Where("group_id = ?", groupID).Find(&groupMenus).Error
	if err != nil {
		return nil, err
	}
	return groupMenus, nil
}

func (s *GroupMenuService) GetAGroupMenus(ctx context.Context, groupID uint) ([]menu.Menu, error) {
	var menus []menu.Menu
	groupMenus, err := s.GetGroupMenus(ctx, groupID)
	if err != nil {
		return nil, err
	}
	for _, groupMenu := range groupMenus {
		if groupMenu.Menu == nil {
			continue
		}
		menus = append(menus, *groupMenu.Menu)
	}

	// Create a copy of the menus slice to sort by the Ordering column
	sortedMenus := make([]menu.Menu, len(menus))
	copy(sortedMenus, menus)
	slices.SortFunc(sortedMenus, func(a, b menu.Menu) int {
		if a.Ordering < b.Ordering {
			return -1
		}
		if a.Ordering > b.Ordering {
			return 1
		}
		return 0
	})

	return sortedMenus, nil
}
