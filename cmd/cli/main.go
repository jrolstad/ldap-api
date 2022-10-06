package main

import (
	"flag"
	"github.com/jrolstad/ldap-api/internal/pkg/configuration"
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/messaging"
	"github.com/jrolstad/ldap-api/internal/pkg/orchestration"
	"github.com/jrolstad/ldap-api/internal/pkg/publishers"
)

var (
	directoryArgument = flag.String("directory", "", "Directory to Search")
	objectArgument    = flag.String("object", "user", "Type of object to search.  Default is user")
)

func main() {
	flag.Parse()

	configurationService := configuration.NewConfigurationService()
	directoryService := directory.NewDirectoryService()
	directoryProcessingServiceFactory := directory.NewDirectoryProcessingServiceFactory(configurationService)
	messageHub := messaging.NewMessageHub(configurationService)
	publisher := publishers.NewDirectoryObjectPublisher(configurationService, messageHub)

	if *objectArgument == "group" {
		orchestration.ProcessAllGroups(*directoryArgument, directoryService, directoryProcessingServiceFactory, publisher)
	} else {
		orchestration.ProcessAllUsers(*directoryArgument, directoryService, directoryProcessingServiceFactory, publisher)
	}
}
