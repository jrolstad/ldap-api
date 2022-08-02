package directory

import (
	"github.com/jrolstad/ldap-api/internal/pkg/configuration"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
)

type DirectorySearchServiceFactory interface {
	NewDirectorySearchService(directory *models.Directory) DirectorySearchService
}

func NewDirectorySearchServiceFactory(configuration configuration.ConfigurationService) DirectorySearchServiceFactory {
	return &directorySearchServiceFactory{
		configurationService: configuration,
	}
}

type directorySearchServiceFactory struct {
	configurationService configuration.ConfigurationService
}

func (s *directorySearchServiceFactory) NewDirectorySearchService(directory *models.Directory) DirectorySearchService {
	ldapUser := s.configurationService.GetValue(directory.UserConfigurationName)
	ldapPassword := s.configurationService.GetSecret(directory.PasswordConfigurationName)

	return &activeDirectorySearchService{
		connection: getLdapConnection(directory.HostName, ldapUser, ldapPassword),
		baseDN:     directory.BaseDN,
	}
}
