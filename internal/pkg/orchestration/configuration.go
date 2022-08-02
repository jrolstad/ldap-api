package orchestration

import (
	"github.com/jrolstad/ldap-api/internal/pkg/models"
	"os"
)

func GetConfiguration() *models.Configuration {
	return &models.Configuration{
		LdapHost:         os.Getenv("ldap_host"),
		LdapUserName:     os.Getenv("ldap_user_name"),
		LdapUserPassword: os.Getenv("ldap_user_password"),
	}
}
