package directory

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
)

type activeDirectorySearchService struct {
	connection *ldap.Conn
	baseDN     string
}

func (s *activeDirectorySearchService) GetUser(alias string) (*models.User, error) {

	filterCriteria := fmt.Sprintf("(&(objectClass=user)(sAMAccountName=%v))", alias)
	fields := getUserAttributes()

	result, searchError := s.searchSingle(filterCriteria, fields)
	if result == nil || searchError != nil {
		return nil, searchError
	}
	user := MapSearchResultToUser(result)
	return user, nil
}

func (s *activeDirectorySearchService) GetUserDirects(alias string) ([]*models.User, error) {

	user, err := s.GetUser(alias)
	if err != nil || user == nil {
		return nil, err
	}

	filterCriteria := fmt.Sprintf("(&(manager=%v))", user.Location)
	fields := getUserAttributes()

	result, searchError := s.search(filterCriteria, fields)
	if result == nil || searchError != nil {
		return nil, searchError
	}
	directs := make([]*models.User, len(result))

	for index, item := range result {
		member := MapSearchResultToUser(item)
		directs[index] = member
	}

	return directs, nil
}

func (s *activeDirectorySearchService) GetGroup(alias string) (*models.Group, error) {
	filterCriteria := fmt.Sprintf("(&(objectClass=group)(sAMAccountName=%v))", alias)
	fields := getGroupAttributes()

	result, searchError := s.searchSingle(filterCriteria, fields)
	if result == nil || searchError != nil {
		return nil, searchError
	}
	return MapSearchResultToGroup(result), nil
}

func (s *activeDirectorySearchService) GetGroupMembers(name string) ([]*models.User, error) {
	group, err := s.GetGroup(name)
	if err != nil || group == nil {
		return nil, err
	}
	filterCriteria := fmt.Sprintf("(memberOf=%v)", group.Location)
	fields := getUserAttributes()

	searchResults, err := s.search(filterCriteria, fields)
	if searchResults == nil || len(searchResults) == 0 {
		return make([]*models.User, 0), err
	}

	members := make([]*models.User, len(searchResults))

	for index, item := range searchResults {
		member := MapSearchResultToUser(item)
		members[index] = member
	}

	return members, nil
}

func (s *activeDirectorySearchService) searchSingle(filter string, fields []string) (*ldap.Entry, error) {
	results, err := s.search(filter, fields)
	if results == nil || len(results) == 0 {
		return nil, err
	}

	return results[0], err
}

func (s *activeDirectorySearchService) search(filter string, fields []string) ([]*ldap.Entry, error) {

	searchRequest := ldap.NewSearchRequest(
		s.baseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter, // The filter to apply
		fields, // A list attributes to retrieve
		nil,
	)

	searchResults, err := s.connection.SearchWithPaging(searchRequest, 100)

	if err != nil {
		return make([]*ldap.Entry, 0), err
	}

	if searchResults == nil {
		return make([]*ldap.Entry, 0), nil
	}

	return searchResults.Entries, nil
}

func (s *activeDirectorySearchService) Close() {
	s.connection.Close()
}
