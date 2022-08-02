package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/orchestration"
)

func configureRoutes(ginHost *gin.Engine, directoryService directory.DirectorySearchService) {
	configureUserRoutes(ginHost, directoryService)
	configureGroupRoutes(ginHost, directoryService)

}

func configureUserRoutes(ginHost *gin.Engine, directoryService directory.DirectorySearchService) {
	ginHost.GET("/:directory/user/:name", func(c *gin.Context) {
		domain := c.Param("directory")
		name := c.Param("name")
		data, err := orchestration.GetUser(domain, name, directoryService)

		returnResult(c, data, err)
	})
}

func configureGroupRoutes(ginHost *gin.Engine, directoryService directory.DirectorySearchService) {
	ginHost.GET("/:directory/group/:name", func(c *gin.Context) {
		domain := c.Param("directory")
		name := c.Param("name")
		data, err := orchestration.GetGroup(domain, name, directoryService)

		returnResult(c, data, err)
	})

	ginHost.GET("/:directory/group/:name/member", func(c *gin.Context) {
		domain := c.Param("directory")
		name := c.Param("name")
		data, err := orchestration.GetGroupMembers(domain, name, directoryService)

		returnResult(c, data, err)
	})
}
