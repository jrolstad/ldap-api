package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jrolstad/ldap-api/internal/pkg/core"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
	"github.com/jrolstad/ldap-api/internal/pkg/repositories"
)

var (
	repository repositories.DirectoryObjectRepository
)

func init() {
	repository = repositories.NewDirectoryObjectRepository()
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.SQSEvent) error {

	for _, item := range event.Records {
		err := processEvent(item)

		if err != nil {
			return err
		}
	}

	return nil
}

func processEvent(message events.SQSMessage) error {
	directoryObject := &models.DirectoryObject{}
	core.MapFromJson(message.Body, directoryObject)
	directoryObject.Data = message.Body

	err := repository.Save(directoryObject)
	return err
}
