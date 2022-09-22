package publishers

import (
	"errors"
	"github.com/jrolstad/ldap-api/internal/pkg/messaging"
	"log"
)

type SqsDirectoryObjectPublisher struct {
	target     string
	messageHub messaging.MessageHub
}

func (s *SqsDirectoryObjectPublisher) Publish(toPublish interface{}) error {
	if toPublish == nil {
		return errors.New("unable to publish nil object")
	}

	err := s.messageHub.Send(toPublish, s.target)

	if err != nil {
		log.Fatalln(err)
	}
	return err

}
