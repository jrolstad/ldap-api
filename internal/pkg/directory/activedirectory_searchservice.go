package directory

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
	"log"
)

type activeDirectorySearchService struct {
	connection *ldap.Conn
}

func (s *activeDirectorySearchService) GetUser(domain string, alias string) (*models.User, error) {

	filterCriteria := fmt.Sprintf("(&(objectClass=user)(sAMAccountName=%v))", alias)
	fields := []string{"objectGUID", "sAMAccountName", "mail", "userPrincipalName", "givenName", "sn", "distinguishedName"}

	result, searchError := s.SearchSingle(filterCriteria, fields)
	if result == nil || searchError != nil {
		return nil, searchError
	}
	user := s.mapSearchResultToUser(result)
	return user, nil
}

func (s *activeDirectorySearchService) GetGroup(domain string, alias string) (*models.Group, error) {
	filterCriteria := fmt.Sprintf("(&(objectClass=group)(sAMAccountName=%v))", alias)
	fields := []string{"objectGUID", "sAMAccountName", "groupType", "distinguishedName"}

	result, searchError := s.SearchSingle(filterCriteria, fields)
	if result == nil || searchError != nil {
		return nil, searchError
	}
	return &models.Group{
		Id:       result.GetAttributeValue("objectGUID"),
		Location: result.GetAttributeValue("distinguishedName"),
		Name:     result.GetAttributeValue("sAMAccountName"),
		Type:     result.GetAttributeValue("groupType"),
	}, nil
}

func (s *activeDirectorySearchService) GetGroupMembers(domain string, name string) ([]*models.User, error) {
	group, err := s.GetGroup(domain, name)
	if err != nil || group == nil {
		return nil, err
	}
	filterCriteria := fmt.Sprintf("(memberOf=%v)", group.Location)
	fields := []string{"objectGUID", "sAMAccountName", "mail", "userPrincipalName", "givenName", "sn", "distinguishedName"}

	searchResults, err := s.Search(filterCriteria, fields)
	if searchResults == nil || len(searchResults) == 0 {
		return make([]*models.User, 0), err
	}

	members := make([]*models.User, len(searchResults))

	for index, item := range searchResults {
		member := s.mapSearchResultToUser(item)
		members[index] = member
	}

	return members, nil
}

func (s *activeDirectorySearchService) mapSearchResultToUser(result *ldap.Entry) *models.User {
	return &models.User{
		Id:        result.GetAttributeValue("objectGUID"),
		Location:  result.GetAttributeValue("distinguishedName"),
		Upn:       result.GetAttributeValue("userPrincipalName"),
		Email:     result.GetAttributeValue("mail"),
		Name:      result.GetAttributeValue("sAMAccountName"),
		GivenName: result.GetAttributeValue("givenName"),
		Surname:   result.GetAttributeValue("sn"),
	}
}

func (s *activeDirectorySearchService) SearchSingle(filter string, fields []string) (*ldap.Entry, error) {
	results, err := s.Search(filter, fields)
	if results == nil || len(results) == 0 {
		return nil, err
	}

	return results[0], err
}

func (s *activeDirectorySearchService) Search(filter string, fields []string) ([]*ldap.Entry, error) {

	searchRequest := ldap.NewSearchRequest(
		"DC=internal,DC=salesforce,DC=com", // The base dn to search
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

func getLdapConnection(host string, userName string, password string) *ldap.Conn {
	address := fmt.Sprintf("ldaps://%v", host)
	conn, err := ldap.DialURL(address)
	if err != nil {
		log.Fatalf("Failed to connect: %s\n", err)
	}

	err = conn.Bind(userName, password)
	if err != nil {
		log.Fatalf("Failed to bind: %s\n", err)
	}

	return conn
}
