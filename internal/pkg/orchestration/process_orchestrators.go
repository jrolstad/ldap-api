package orchestration

import (
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
	"github.com/jrolstad/ldap-api/internal/pkg/publishers"
)

func ProcessAllUsers(directoryName string, directoryService *directory.DirectoryService,
	processingServiceFactory directory.DirectoryProcessingServiceFactory,
	publisher publishers.DirectoryObjectPublisher) error {
	directory, err := directoryService.Get(directoryName)
	if err != nil || directory == nil {
		return err
	}

	processingService := processingServiceFactory.NewDirectoryProcessingService(directory)
	defer processingService.Close()

	processor := func(data []*models.User) {
		for _, item := range data {
			publisher.Publish(item)
		}
	}
	return processingService.ProcessAllUsers(processor)
}

func ProcessAllGroups(directoryName string, directoryService *directory.DirectoryService,
	processingServiceFactory directory.DirectoryProcessingServiceFactory,
	publisher publishers.DirectoryObjectPublisher) error {
	directory, err := directoryService.Get(directoryName)
	if err != nil || directory == nil {
		return err
	}

	processingService := processingServiceFactory.NewDirectoryProcessingService(directory)
	defer processingService.Close()

	processor := func(data []*models.Group) {
		for _, item := range data {
			publisher.Publish(item)
		}
	}
	return processingService.ProcessAllGroups(processor)
}
