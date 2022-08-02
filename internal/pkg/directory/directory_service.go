package directory

import "github.com/jrolstad/ldap-api/internal/pkg/models"

type DirectoryService interface {
	GetUser(domain string, alias string) *models.User
	GetSecurityGroup(domain string, alias string) *models.Group
}

func NewDirectoryService() DirectoryService {
	return &activeDirectoryService{}
}
