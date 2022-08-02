package directory

import "github.com/jrolstad/ldap-api/internal/pkg/models"

type DirectoryService interface {
	GetUser(domain string, alias string) (*models.User, error)
	GetSecurityGroup(domain string, alias string) (*models.Group, error)
	Close()
}

func NewDirectoryService(config *models.Configuration) DirectoryService {
	return &activeDirectoryService{
		connection: getLdapConnection(config.LdapHost, config.LdapUserName, config.LdapUserPassword),
	}
}
