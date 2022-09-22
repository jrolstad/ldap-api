package publishers

import "github.com/jrolstad/ldap-api/internal/pkg/configuration"

type DirectoryObjectPublisher interface {
	Publish(toPublish interface{}) error
}

func NewDirectoryObjectPublisher(configuration *configuration.ConfigurationService) DirectoryObjectPublisher {
	instance := &SqsDirectoryObjectPublisher{configuration: configuration}

	return instance
}
