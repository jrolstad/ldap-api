package directory

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
	"log"
)

type activeDirectoryService struct {
	connection *ldap.Conn
}

func (s *activeDirectoryService) GetUser(domain string, alias string) (*models.User, error) {

	filterCriteria := fmt.Sprintf("(&(objectClass=user)(sAMAccountName=%v))", alias)
	fields := []string{"objectGUID", "sAMAccountName", "mail", "userPrincipalName", "givenName", "sn"}

	result, searchError := s.SearchSingle(filterCriteria, fields)
	if result == nil || searchError != nil {
		return nil, searchError
	}
	return &models.User{
		Id:        result.GetAttributeValue("objectGUID"),
		Upn:       result.GetAttributeValue("userPrincipalName"),
		Email:     result.GetAttributeValue("mail"),
		Name:      result.GetAttributeValue("sAMAccountName"),
		GivenName: result.GetAttributeValue("givenName"),
		Surname:   result.GetAttributeValue("sn"),
	}, nil
}

func (s *activeDirectoryService) GetSecurityGroup(domain string, alias string) (*models.Group, error) {
	return &models.Group{
		Id:      "{2f3e225a-5fff-4049-8590-d3e6a96aff09}",
		Domain:  domain,
		Name:    "BI_Alliances_Channels_Project_Leaders",
		Members: make([]*models.User, 0),
	}, nil
}

func (s *activeDirectoryService) SearchSingle(filter string, fields []string) (*ldap.Entry, error) {
	results, err := s.Search(filter, fields)
	if results == nil || len(results) == 0 {
		return nil, err
	}

	return results[0], err
}

func (s *activeDirectoryService) Search(filter string, fields []string) ([]*ldap.Entry, error) {

	searchRequest := ldap.NewSearchRequest(
		"DC=internal,DC=salesforce,DC=com", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter, // The filter to apply
		fields, // A list attributes to retrieve
		nil,
	)

	searchResults, err := s.connection.Search(searchRequest)
	if err != nil {
		return make([]*ldap.Entry, 0), err
	}

	if searchResults == nil {
		return make([]*ldap.Entry, 0), nil
	}

	return searchResults.Entries, nil
}

func (s *activeDirectoryService) Close() {
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
