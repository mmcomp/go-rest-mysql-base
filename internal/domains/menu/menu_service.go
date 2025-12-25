package menu

import (
	"slices"

	"gorm.io/gorm"
)

type MenuService struct {
	db *gorm.DB
}

func NewMenuService(db *gorm.DB) *MenuService {
	return &MenuService{db: db}
}

type MenuNode struct {
	ID       uint       `json:"id"`
	Name     string     `json:"name"`
	Path     string     `json:"path"`
	ParentID uint       `json:"-"`
	Children []MenuNode `json:"children"`
	Order    int        `json:"-"`
}

func loadMenusRecursively(nodes []MenuNode, child *MenuNode) {
	for j := range nodes {
		if nodes[j].ID == child.ParentID {
			nodes[j].Children = append(nodes[j].Children, *child)
			continue
		}
		loadMenusRecursively(nodes[j].Children, child)
	}
}

func (s *MenuService) ArrangeMenusTreeLikeLike(menus []Menu) []MenuNode {
	menuToNode := func(menu Menu) MenuNode {
		return MenuNode{
			ID:       menu.ID,
			Name:     menu.Name,
			Path:     menu.Path,
			ParentID: menu.ParentID,
			Children: make([]MenuNode, 0),
			Order:    menu.Ordering,
		}
	}
	menusMap := make(map[uint]*MenuNode)

	for _, menu := range menus {
		m := menuToNode(menu)
		if p, exists := menusMap[menu.ParentID]; exists {
			p.Children = append(p.Children, m)
		} else {
			menusMap[menu.ID] = &m
		}
	}

	for k, menu := range menusMap {
		if menu.ParentID == 0 {
			continue
		}
		if p, exists := menusMap[menu.ParentID]; exists {
			p.Children = append(p.Children, *menu)
		} else {
			for _, mp := range menusMap {
				loadMenusRecursively(mp.Children, menu)
			}
		}
		delete(menusMap, k)
	}

	menuResult := make([]MenuNode, 0, len(menusMap))
	for _, menu := range menusMap {
		menuResult = append(menuResult, *menu)
	}
	// Copy menuResult to avoid mutating the original slice
	sortedMenuResult := make([]MenuNode, len(menuResult))
	copy(sortedMenuResult, menuResult)

	// Sort using slices.SortFunc, passing values instead of pointers
	slices.SortFunc(sortedMenuResult, func(a, b MenuNode) int {
		if a.Order < b.Order {
			return -1
		}
		if a.Order > b.Order {
			return 1
		}
		return 0
	})
	return sortedMenuResult
}
