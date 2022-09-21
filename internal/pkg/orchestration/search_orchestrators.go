package orchestration

import (
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
)

func GetUser(directoryName string, name string, directoryService *directory.DirectoryService, searchServiceFactory directory.DirectorySearchServiceFactory) (*models.User, error) {
	directory, err := directoryService.Get(directoryName)
	if err != nil || directory == nil {
		return nil, err
	}

	searchService := searchServiceFactory.NewDirectorySearchService(directory)
	defer searchService.Close()

	return searchService.GetUser(name)
}

func GetUserSubordinates(directoryName string, name string, directoryService *directory.DirectoryService, searchServiceFactory directory.DirectorySearchServiceFactory) ([]*models.User, error) {
	directory, err := directoryService.Get(directoryName)
	if err != nil || directory == nil {
		return nil, err
	}

	searchService := searchServiceFactory.NewDirectorySearchService(directory)
	defer searchService.Close()

	return searchService.GetUserSubordinates(name)
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
