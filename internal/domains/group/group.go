package group

import "time"

type Group struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Type      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (Group) TableName() string {
	return "groups"
}
