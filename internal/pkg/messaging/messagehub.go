package messaging

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/jrolstad/ldap-api/internal/pkg/configuration"
	"github.com/jrolstad/ldap-api/internal/pkg/core"
)

type MessageHub interface {
	Send(toSend interface{}, target string) error
}

func NewMessageHub(config configuration.ConfigurationService) MessageHub {
	hub := new(SqsMessageHub)

	session := core.GetAwsSession(config)
	hub.sqs = sqs.New(session)

	return hub
}
