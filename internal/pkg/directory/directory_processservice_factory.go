package directory

import (
	"github.com/jrolstad/ldap-api/internal/pkg/configuration"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
)

type DirectoryProcessingServiceFactory interface {
	NewDirectoryProcessingService(directory *models.Directory) DirectoryProcessingService
}

func NewDirectoryProcessingServiceFactory(configuration configuration.ConfigurationService) DirectoryProcessingServiceFactory {
	return &directoryProcessingServiceFactory{
		configurationService: configuration,
	}
}

type directoryProcessingServiceFactory struct {
	configurationService configuration.ConfigurationService
}

func (s *directoryProcessingServiceFactory) NewDirectoryProcessingService(directory *models.Directory) DirectoryProcessingService {
	ldapUser := s.configurationService.GetValue(directory.UserConfigurationName)
	ldapPassword := s.configurationService.GetSecret(directory.PasswordConfigurationName)

	return &activeDirectoryProcessingService{
		connection: getLdapConnection(directory.HostName, ldapUser, ldapPassword),
		baseDN:     directory.BaseDN,
	}
}
