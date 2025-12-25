package routes

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mmcomp/go-rest-mysql-base/docs"
	"github.com/mmcomp/go-rest-mysql-base/internal/domains/auth"
	"github.com/mmcomp/go-rest-mysql-base/middlewares"
	"github.com/mmcomp/go-rest-mysql-base/ratelimit"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, secretKey string, db *gorm.DB) {
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 50,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: func(c *gin.Context, info ratelimit.Info) {
			c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
		},
		KeyFunc: func(c *gin.Context) string {
			return c.ClientIP()
		},
	})
	router.Static("/static", "./static")

	router.Use(mw)

	unauthenticated := router.Group("/")
	authorized := router.Group("/")
	authorized.Use(middlewares.AuthMiddleware(secretKey))
	auth.InitAuthDomain(authorized, unauthenticated, secretKey, db)
	if os.Getenv("ENV") != "production" {
		docs.SwaggerInfo.BasePath = "/api/v1"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}
