package main

import (
	"flag"
	"github.com/jrolstad/ldap-api/internal/pkg/configuration"
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/orchestration"
	"github.com/jrolstad/ldap-api/internal/pkg/publishers"
)

var (
	directoryArgument = flag.String("directory", "", "Directory to Search")
)

func main() {
	flag.Parse()

	configurationService := configuration.NewConfigurationService()
	directoryService := directory.NewDirectoryService()
	directoryProcessingServiceFactory := directory.NewDirectoryProcessingServiceFactory(configurationService)
	publisher := publishers.NewDirectoryObjectPublisher(&configurationService)

	orchestration.ProcessAllUsers(*directoryArgument, directoryService, directoryProcessingServiceFactory, publisher)
}
