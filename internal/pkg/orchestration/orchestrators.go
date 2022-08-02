package orchestration

import (
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
)

func GetUser(domain string, alias string, service directory.DirectorySearchService) (*models.User, error) {
	return service.GetUser(domain, alias)
}

func GetGroup(domain string, alias string, service directory.DirectorySearchService) (*models.Group, error) {
	return service.GetGroup(domain, alias)
}
