package directory

import "github.com/jrolstad/ldap-api/internal/pkg/models"

type activeDirectoryService struct {
}

func (s *activeDirectoryService) GetUser(domain string, alias string) *models.User {
	return &models.User{
		Id:        "{8c7247f7-def1-4373-8bb9-3ce8b40117ff}",
		Upn:       "jrolstad@internal.salesforce.com",
		Email:     "jrolstad@salesforce.com",
		GivenName: "Josh",
		Surname:   "Rolstad",
	}
}
