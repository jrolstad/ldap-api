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
	ginHost.GET("/user/:domain/:alias", func(c *gin.Context) {
		domain := c.Param("domain")
		alias := c.Param("alias")
		data, err := orchestration.GetUser(domain, alias, directoryService)

		returnResult(c, data, err)
	})
}

func configureGroupRoutes(ginHost *gin.Engine, directoryService directory.DirectorySearchService) {
	ginHost.GET("/group/:domain/:alias", func(c *gin.Context) {
		domain := c.Param("domain")
		alias := c.Param("alias")
		data, err := orchestration.GetGroup(domain, alias, directoryService)

		returnResult(c, data, err)
	})
}
