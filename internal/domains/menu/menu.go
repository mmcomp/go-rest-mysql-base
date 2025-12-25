package menu

import "time"

type Menu struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Path      string    `gorm:"not null"`
	ParentID  uint      `gorm:"not null"`
	Ordering  int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (Menu) TableName() string {
	return "menus"
}
