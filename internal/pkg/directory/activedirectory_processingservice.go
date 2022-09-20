package directory

import (
	"github.com/go-ldap/ldap/v3"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
)

type activeDirectoryProcessingService struct {
	connection *ldap.Conn
	baseDN     string
}

func (s *activeDirectoryProcessingService) ProcessUsers(action func([]*models.User)) error {
	return nil
}

func (s *activeDirectoryProcessingService) Close() {
	s.connection.Close()
}
