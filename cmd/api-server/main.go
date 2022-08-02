package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/orchestration"
)

func main() {
	directoryService := directory.NewDirectoryService()
	r := gin.Default()

	r.GET("/user/:domain/*alias", func(c *gin.Context) {
		domain := c.Param("domain")
		alias := c.Param("alias")
		data := orchestration.GetUser(domain, alias, directoryService)

		c.JSON(200, data)
	})

	r.GET("/group/security/:domain/*alias", func(c *gin.Context) {
		domain := c.Param("domain")
		alias := c.Param("alias")
		data := orchestration.GetSecurityGroup(domain, alias, directoryService)

		c.JSON(200, data)
	})

	r.Run(":8080")
}
