package database

import (
	"fmt"

	"github.com/mmcomp/go-rest-mysql-base/internal/domains/auth"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(host, port, user, password, dbname string, config *gorm.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	return gorm.Open(mysql.Open(dsn), config)
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&auth.User{},
	)
}
