package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/orchestration"
)

func main() {
	config := orchestration.GetConfiguration()
	directoryService := directory.NewDirectoryService(config)
	defer directoryService.Close()

	ginHost := gin.Default()

	configureRoutes(ginHost, directoryService)
	ginHost.Run(":8080")
}

func configureRoutes(ginHost *gin.Engine, directoryService directory.DirectoryService) {

	ginHost.GET("/user/:domain/:alias", func(c *gin.Context) {
		domain := c.Param("domain")
		alias := c.Param("alias")
		data, err := orchestration.GetUser(domain, alias, directoryService)

		returnResult(c, data, err)
	})

	ginHost.GET("/group/:domain/:alias", func(c *gin.Context) {
		domain := c.Param("domain")
		alias := c.Param("alias")
		data, err := orchestration.GetGroup(domain, alias, directoryService)

		returnResult(c, data, err)
	})
}

func returnResult(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(500, "Error when processing request")
	} else {
		c.JSON(200, data)
	}
}
