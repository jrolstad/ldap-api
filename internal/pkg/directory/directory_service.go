package directory

import "github.com/jrolstad/ldap-api/internal/pkg/models"

type DirectoryService interface {
	GetUser(domain string, alias string) *models.User
}

func NewDirectoryService() DirectoryService {
	return &activeDirectoryService{}
}
