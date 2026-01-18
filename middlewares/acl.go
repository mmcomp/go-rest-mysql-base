package middlewares

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/mmcomp/go-rest-mysql-base/internal/domains/menu"
)

var GroupMenus map[uint][]menu.Menu
var freePaths []string = []string{"/api/v1/auth/login", "/api/v1/auth/refresh-token", "/api/v1/auth/ping"}

func ACLMiddleware(c *gin.Context) {
	url := c.Request.URL.Path
	groupID := c.GetUint("group")
	menus, ok := GroupMenus[groupID]
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}
	for _, menu := range menus {
		if "/api/v1/"+menu.Path == url || slices.Contains(freePaths, url) {
			c.Next()
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	c.Abort()
}
