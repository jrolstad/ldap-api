package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/orchestration"
)

func configureRoutes(ginHost *gin.Engine,
	directoryService *directory.DirectoryService,
	searchServiceFactory directory.DirectorySearchServiceFactory) {
	configureUserRoutes(ginHost, directoryService, searchServiceFactory)
	configureGroupRoutes(ginHost, directoryService, searchServiceFactory)
}

func configureUserRoutes(ginHost *gin.Engine,
	directoryService *directory.DirectoryService,
	searchServiceFactory directory.DirectorySearchServiceFactory) {
	ginHost.GET("/:directory/user/:name", func(c *gin.Context) {
		directory := c.Param("directory")
		name := c.Param("name")
		data, err := orchestration.GetUser(directory, name, directoryService, searchServiceFactory)

		returnJsonResult(c, data, err)
	})
	ginHost.GET("/:directory/user/:name/directs", func(c *gin.Context) {
		directory := c.Param("directory")
		name := c.Param("name")
		data, err := orchestration.GetUserDirects(directory, name, directoryService, searchServiceFactory)

		returnJsonResult(c, data, err)
	})
}

func configureGroupRoutes(ginHost *gin.Engine,
	directoryService *directory.DirectoryService,
	searchServiceFactory directory.DirectorySearchServiceFactory) {

	ginHost.GET("/:directory/group/:name", func(c *gin.Context) {
		directory := c.Param("directory")
		name := c.Param("name")
		data, err := orchestration.GetGroup(directory, name, directoryService, searchServiceFactory)

		returnJsonResult(c, data, err)
	})

	ginHost.GET("/:directory/group/:name/member", func(c *gin.Context) {
		directory := c.Param("directory")
		name := c.Param("name")
		data, err := orchestration.GetGroupMembers(directory, name, directoryService, searchServiceFactory)

		returnJsonResult(c, data, err)
	})
}
