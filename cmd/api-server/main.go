package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/orchestration"
)

func main() {
	ginHost := gin.Default()

	configureRoutes(ginHost)
	ginHost.Run(":8080")
}

func configureRoutes(ginHost *gin.Engine) {
	config := orchestration.GetConfiguration()
	directoryService := directory.NewDirectoryService(config)

	ginHost.GET("/user/:domain/*alias", func(c *gin.Context) {
		domain := c.Param("domain")
		alias := c.Param("alias")
		data, err := orchestration.GetUser(domain, alias, directoryService)

		if err != nil {
			c.JSON(500, "Error when retrieving user")
		} else {
			c.JSON(200, data)
		}
	})

	ginHost.GET("/group/security/:domain/*alias", func(c *gin.Context) {
		domain := c.Param("domain")
		alias := c.Param("alias")
		data, err := orchestration.GetSecurityGroup(domain, alias, directoryService)

		if err != nil {
			c.JSON(500, "Error when retrieving group")
		} else {
			c.JSON(200, data)
		}
	})
}
