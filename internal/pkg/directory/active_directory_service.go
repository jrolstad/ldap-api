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
	searchRequest := ldap.NewSearchRequest(
		"DC=internal,DC=salesforce,DC=com", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filterCriteria, // The filter to apply
		[]string{"objectGUID", "sAMAccountName", "mail", "userPrincipalName", "givenName", "sn"}, // A list attributes to retrieve
		nil,
	)

	searchResults, err := s.connection.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	if searchResults == nil || len(searchResults.Entries) == 0 {
		return nil, nil
	}

	data := searchResults.Entries[0]
	return &models.User{
		Id:        data.GetAttributeValue("objectGUID"),
		Upn:       data.GetAttributeValue("userPrincipalName"),
		Email:     data.GetAttributeValue("mail"),
		Name:      data.GetAttributeValue("sAMAccountName"),
		GivenName: data.GetAttributeValue("givenName"),
		Surname:   data.GetAttributeValue("sn"),
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
