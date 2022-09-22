package publishers

import (
	"errors"
	"github.com/jrolstad/ldap-api/internal/pkg/configuration"
	"github.com/jrolstad/ldap-api/internal/pkg/core"
	"log"
)

type SqsDirectoryObjectPublisher struct {
	configuration *configuration.ConfigurationService
}

func (s *SqsDirectoryObjectPublisher) Publish(toPublish interface{}) error {
	if toPublish == nil {
		return errors.New("unable to publish nil object")
	}

	serializedObject := core.MapToJson(toPublish)
	log.Println(serializedObject)

	return nil

}
