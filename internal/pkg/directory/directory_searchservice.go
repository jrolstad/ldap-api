package directory

import "github.com/jrolstad/ldap-api/internal/pkg/models"

type DirectorySearchService interface {
	GetUsers() ([]*models.User, error)
	GetUser(name string) (*models.User, error)
	GetGroup(name string) (*models.Group, error)
	GetGroupMembers(name string) ([]*models.User, error)
	Close()
}

type DirectoryProcessingService interface {
	ProcessUsers(action func([]*models.User)) error
}
