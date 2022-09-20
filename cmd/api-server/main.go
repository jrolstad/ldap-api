package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jrolstad/ldap-api/internal/pkg/configuration"
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
)

func main() {
	configurationService := configuration.NewConfigurationService()
	directoryService := directory.NewDirectoryService()
	directorySearchServiceFactory := directory.NewDirectorySearchServiceFactory(configurationService)
	directoryProcessingServiceFactory := directory.NewDirectoryProcessingServiceFactory(configurationService)
	
	ginHost := gin.Default()

	configureRoutes(ginHost, directoryService, directorySearchServiceFactory, directoryProcessingServiceFactory)
	ginHost.Run(":8080")
}
