package directory

import (
	"github.com/jrolstad/ldap-api/internal/pkg/models"
	"strings"
)

type DirectoryService struct {
}

func NewDirectoryService() *DirectoryService {
	return &DirectoryService{}
}

func (s *DirectoryService) Get(name string) (*models.Directory, error) {
	if strings.EqualFold(name, "internal.salesforce.com") {
		return &models.Directory{
			Id:                        "sfdc-internal",
			Name:                      "internal.salesforce.com",
			BaseDN:                    "DC=internal,DC=salesforce,DC=com",
			HostName:                  "internal.salesforce.com:636",
			UserConfigurationName:     "ldap_user_name",
			PasswordConfigurationName: "ldap_user_password",
		}, nil
	}

	return nil, nil
}
