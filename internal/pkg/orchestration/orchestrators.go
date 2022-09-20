package orchestration

import (
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
)

func ProcessUsers(directoryName string, directoryService *directory.DirectoryService, processingServiceFactory directory.DirectoryProcessingServiceFactory) error {
	directory, err := directoryService.Get(directoryName)
	if err != nil || directory == nil {
		return err
	}

	processingService := processingServiceFactory.NewDirectoryProcessingService(directory)
	defer processingService.Close()

	return processingService.ProcessUsers(nil)
}

func GetUser(directoryName string, name string, directoryService *directory.DirectoryService, searchServiceFactory directory.DirectorySearchServiceFactory) (*models.User, error) {
	directory, err := directoryService.Get(directoryName)
	if err != nil || directory == nil {
		return nil, err
	}

	searchService := searchServiceFactory.NewDirectorySearchService(directory)
	defer searchService.Close()

	return searchService.GetUser(name)
}

func GetGroup(directoryName string, name string, directoryService *directory.DirectoryService, searchServiceFactory directory.DirectorySearchServiceFactory) (*models.Group, error) {
	directory, err := directoryService.Get(directoryName)
	if err != nil || directory == nil {
		return nil, err
	}

	searchService := searchServiceFactory.NewDirectorySearchService(directory)
	defer searchService.Close()

	return searchService.GetGroup(name)
}

func GetGroupMembers(directoryName string, name string, directoryService *directory.DirectoryService, searchServiceFactory directory.DirectorySearchServiceFactory) ([]*models.User, error) {
	directory, err := directoryService.Get(directoryName)
	if err != nil || directory == nil {
		return nil, err
	}

	searchService := searchServiceFactory.NewDirectorySearchService(directory)
	defer searchService.Close()

	return searchService.GetGroupMembers(name)
}
