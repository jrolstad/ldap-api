package publishers

import (
	"github.com/jrolstad/ldap-api/internal/pkg/configuration"
	"github.com/jrolstad/ldap-api/internal/pkg/messaging"
)

type DirectoryObjectPublisher interface {
	Publish(toPublish interface{}) error
}

func NewDirectoryObjectPublisher(configuration configuration.ConfigurationService, messageHub messaging.MessageHub) DirectoryObjectPublisher {
	instance := &SqsDirectoryObjectPublisher{
		target:     configuration.GetValue("directoryobject_queue"),
		messageHub: messageHub,
	}

	return instance
}
