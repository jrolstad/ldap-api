package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
)

func main() {
	directoryService := directory.NewDirectoryService()
	directorySearchServiceFactory := directory.NewDirectorySearchServiceFactory()

	ginHost := gin.Default()

	configureRoutes(ginHost, directoryService, directorySearchServiceFactory)
	ginHost.Run(":8080")
}
