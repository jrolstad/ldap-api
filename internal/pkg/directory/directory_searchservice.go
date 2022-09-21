package directory

import "github.com/jrolstad/ldap-api/internal/pkg/models"

type DirectorySearchService interface {
	GetUser(name string) (*models.User, error)
	GetUserDirects(alias string) ([]*models.User, error)
	GetGroup(name string) (*models.Group, error)
	GetGroupMembers(name string) ([]*models.User, error)
	Close()
}
