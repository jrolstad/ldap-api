package directory

import (
	"errors"
	"github.com/go-ldap/ldap/v3"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
)

type activeDirectoryProcessingService struct {
	connection *ldap.Conn
	baseDN     string
}

func (s *activeDirectoryProcessingService) ProcessAllUsers(action func([]*models.User)) error {
	if action == nil {
		return errors.New("no action defined")
	}

	filterCriteria := "(&(objectClass=user))"
	fields := []string{"objectGUID", "sAMAccountName", "mail", "userPrincipalName", "givenName", "sn", "distinguishedName"}

	processor := func(items []*ldap.Entry) {
		data := make([]*models.User, 0)
		for _, item := range items {
			user := MapSearchResultToUser(item)
			data = append(data, user)
		}
		action(data)
	}

	return s.searchWithAction(filterCriteria, fields, processor)
}

func (s *activeDirectoryProcessingService) searchWithAction(filter string, fields []string, action func([]*ldap.Entry)) error {

	pagingControl := &ldap.ControlPaging{PagingSize: 100}
	searchRequest := ldap.NewSearchRequest(
		s.baseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter, // The filter to apply
		fields, // A list attributes to retrieve
		[]ldap.Control{pagingControl},
	)

	for {
		result, err := s.connection.Search(searchRequest)
		if err != nil {
			return err
		}
		if result == nil {
			return ldap.NewError(ldap.ErrorNetwork, errors.New("ldap: packet not received"))
		}

		action(result.Entries)

		pagingResult := ldap.FindControl(result.Controls, ldap.ControlTypePaging)
		if pagingResult == nil {
			pagingControl = nil
			break
		}

		cookie := pagingResult.(*ldap.ControlPaging).Cookie
		if len(cookie) == 0 {
			pagingControl = nil
			break
		}
		pagingControl.SetCookie(cookie)
	}

	return nil
}

func (s *activeDirectoryProcessingService) Close() {
	s.connection.Close()
}
