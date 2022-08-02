package orchestration

import (
	"github.com/jrolstad/ldap-api/internal/pkg/directory"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
)

func GetUser(domain string, alias string, service directory.DirectoryService) *models.User {
	return service.GetUser(domain, alias)
}
