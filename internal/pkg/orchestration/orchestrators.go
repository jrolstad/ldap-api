package orchestration

import (
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
)

func GetUser(domain string, alias string, service directory.DirectoryService) (*models.User, error) {
	return service.GetUser(domain, alias)
}

func GetGroup(domain string, alias string, service directory.DirectoryService) (*models.Group, error) {
	return service.GetGroup(domain, alias)
}
