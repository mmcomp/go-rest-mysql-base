package user

import (
	"time"

	"github.com/mmcomp/go-rest-mysql-base/internal/domains/group"
)

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

type User struct {
	ID        uint         `gorm:"primaryKey" json:"id"`
	Username  string       `json:"username" gorm:"column:email"`
	Password  string       `json:"-" gorm:"column:password"`
	GroupId   uint         `json:"group_id" gorm:"column:group_id"`
	Group     *group.Group `json:"group" gorm:"foreignKey:GroupId"`
	FirstName string       `json:"first_name" gorm:"column:first_name"`
	LastName  string       `json:"last_name" gorm:"column:last_name"`
	CreatedAt time.Time    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}
