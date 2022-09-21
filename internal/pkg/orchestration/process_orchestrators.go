package orchestration

import (
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
	"log"
)

func ProcessAllUsers(directoryName string, directoryService *directory.DirectoryService, processingServiceFactory directory.DirectoryProcessingServiceFactory) error {
	directory, err := directoryService.Get(directoryName)
	if err != nil || directory == nil {
		return err
	}

	processingService := processingServiceFactory.NewDirectoryProcessingService(directory)
	defer processingService.Close()

	processor := func(data []*models.User) {
		for _, item := range data {
			log.Println(item.Name)
		}
	}
	return processingService.ProcessAllUsers(processor)
}

func ProcessAllGroups(directoryName string, directoryService *directory.DirectoryService, processingServiceFactory directory.DirectoryProcessingServiceFactory) error {
	directory, err := directoryService.Get(directoryName)
	if err != nil || directory == nil {
		return err
	}

	processingService := processingServiceFactory.NewDirectoryProcessingService(directory)
	defer processingService.Close()

	processor := func(data []*models.Group) {
		for _, item := range data {
			log.Println(item.Name)
		}
	}
	return processingService.ProcessAllGroups(processor)
}
