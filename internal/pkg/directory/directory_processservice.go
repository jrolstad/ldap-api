package directory

import "github.com/jrolstad/ldap-api/internal/pkg/models"

type DirectoryProcessingService interface {
	ProcessUsers(action func([]*models.User)) error
	Close()
}
