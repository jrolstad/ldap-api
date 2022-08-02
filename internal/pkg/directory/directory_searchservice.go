package directory

import "github.com/jrolstad/ldap-api/internal/pkg/models"

type DirectorySearchService interface {
	GetUser(domain string, alias string) (*models.User, error)
	GetGroup(domain string, alias string) (*models.Group, error)
	Close()
}

func NewDirectorySearchService(config *models.Configuration) DirectorySearchService {
	return &activeDirectorySearchService{
		connection: getLdapConnection(config.LdapHost, config.LdapUserName, config.LdapUserPassword),
	}
}
