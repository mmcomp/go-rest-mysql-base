package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mmcomp/go-rest-mysql-base/database"
	"github.com/mmcomp/go-rest-mysql-base/routes"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var version = ""

// @securityDefinitions.apikey	AuthToken
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and JWT token.
//
// @securityDefinitions.apikey	RefreshToken
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and JWT refresh token.
func main() {
	// for _, arg := range os.Args {
	// 	fmt.Println("arg: '", arg, "' = migrate:", arg == "migrate")
	// }
	// return
	if len(os.Args) >= 2 && os.Args[1] == "version" {
		fmt.Println(version)
		return
	}

	godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	secretKey := os.Getenv("AUTH_SECRET_KEY")
	appPort := os.Getenv("PORT")
	env := os.Getenv("ENV")
	if env == "" {
		env = "production"
	}
	if appPort == "" {
		appPort = "8080"
	}

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	if env == "production" {
		config = &gorm.Config{}
	}

	db, err := database.Connect(host, port, user, password, dbname, config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	url := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			cmdMigrate(url, os.Args[2:])
			return
		}
	} else if env == "production" {
		cmdMigrate(url, []string{"run"})
	}

	// Init cron
	c := cron.New()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.SetupRoutes(r, secretKey, db)

	c.Start()

	r.Run(":" + appPort)
	fmt.Println("Server is running on port " + appPort)
}
