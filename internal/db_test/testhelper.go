package db_test

import (
	"log"

	"github.com/mmcomp/go-rest-mysql-base/internal/domains/group"
	groupmenu "github.com/mmcomp/go-rest-mysql-base/internal/domains/group_menu"
	"github.com/mmcomp/go-rest-mysql-base/internal/domains/menu"
	"github.com/mmcomp/go-rest-mysql-base/internal/domains/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func PrepareDatabase() (*gorm.DB, func()) {
	database, _ := gorm.Open(sqlite.Open("file:memdb1?mode=memory&cache=shared"), &gorm.Config{})
	models := []any{
		user.User{},
		group.Group{},
		menu.Menu{},
		groupmenu.GroupMenu{},
	}
	for _, model := range models {
		if err := database.AutoMigrate(model); err != nil {
			log.Fatal(err)
		}
	}
	cleanupDB := func() {
		for _, model := range models {
			database.Where("1 = 1").Delete(model)
		}
	}

	return database, cleanupDB
}
