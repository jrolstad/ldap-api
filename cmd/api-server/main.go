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
