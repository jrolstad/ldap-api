package directory

import (
	"github.com/jrolstad/ldap-api/internal/pkg/models"
	"os"
)

type DirectorySearchServiceFactory interface {
	NewDirectorySearchService(directory *models.Directory) DirectorySearchService
}

func NewDirectorySearchServiceFactory() DirectorySearchServiceFactory {
	return &directorySearchServiceFactory{}
}

type directorySearchServiceFactory struct {
}

func (s *directorySearchServiceFactory) NewDirectorySearchService(directory *models.Directory) DirectorySearchService {
	ldapUser := os.Getenv(directory.UserConfigurationName)
	ldapPassword := os.Getenv(directory.PasswordConfigurationName)

	return &activeDirectorySearchService{
		connection: getLdapConnection(directory.HostName, ldapUser, ldapPassword),
		baseDN:     directory.BaseDN,
	}
}
