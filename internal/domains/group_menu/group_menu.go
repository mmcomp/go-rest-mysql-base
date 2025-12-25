package groupmenu

import (
	"time"

	"github.com/mmcomp/go-rest-mysql-base/internal/domains/group"
	"github.com/mmcomp/go-rest-mysql-base/internal/domains/menu"
)

type GroupMenu struct {
	ID        uint         `gorm:"primaryKey"`
	GroupID   uint         `gorm:"not null"`
	MenuID    uint         `gorm:"not null"`
	Menu      *menu.Menu   `gorm:"foreignKey:MenuID"`
	Group     *group.Group `gorm:"foreignKey:GroupID"`
	CreatedAt time.Time    `gorm:"not null"`
	UpdatedAt time.Time    `gorm:"not null"`
}

func (GroupMenu) TableName() string {
	return "group_menus"
}
