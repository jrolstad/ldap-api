package orchestration

import (
	"github.com/jrolstad/ldap-api/internal/pkg/core"
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
	"github.com/jrolstad/ldap-api/internal/pkg/publishers"
	"log"
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
		publishedCount := len(data)
		toPublish := core.ToInterfaceSlice(data)
		publisher.Publish(toPublish)

		log.Printf("Published %v users", publishedCount)
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
		publishedCount := len(data)
		toPublish := core.ToInterfaceSlice(data)
		publisher.Publish(toPublish)
		log.Printf("Published %v groups", publishedCount)
	}
	return processingService.ProcessAllGroups(processor)
}
