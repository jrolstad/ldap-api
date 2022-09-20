package directory

import "github.com/jrolstad/ldap-api/internal/pkg/models"

type DirectoryProcessingService interface {
	ProcessAllUsers(action func([]*models.User)) error
	ProcessAllGroups(action func([]*models.Group)) error
	Close()
}
