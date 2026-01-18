package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func cmdMigrate(url string, args []string) {
	wd, _ := os.Getwd()
	source := fmt.Sprintf("file://%s/migrations/", wd)
	m, err := migrate.New(
		source,
		url)
	if err != nil {
		println("[E] run in the root folder ( backend )")
		log.Fatalf("%v, connection string: %s", err, url)
	}

	if len(args) == 0 {
		log.Fatal("migration command required: {run, force}")
	}

	switch args[0] {
	case "run":
		cmdMigrateRun(m)
		return
	case "down":
		cmdMigrateDown(m)
		return
	case "force":
		cmdMigrateForce(m, args[1:])
		return
	}
}

func cmdMigrateRun(m *migrate.Migrate) {
	if err := m.Up(); err == migrate.ErrNoChange {
		log.Printf("no change")
	} else if err != nil {
		log.Fatal(err)
	}
}

func cmdMigrateForce(m *migrate.Migrate, args []string) {
	if len(args) == 0 {
		log.Fatal("migration force version required: {1, 2, 3, ...}")
	}

	v, err := strconv.Atoi(args[0])
	if err != nil || v <= 0 {
		log.Fatal("migration force version required: {1, 2, 3, ...}")
	}

	if err := m.Force(v); err != nil {
		log.Fatal(err)
	}
}

func cmdMigrateDown(m *migrate.Migrate) {
	if err := m.Down(); err == migrate.ErrNoChange {
		log.Printf("no change")
	} else if err != nil {
		log.Fatal(err)
	}
}
