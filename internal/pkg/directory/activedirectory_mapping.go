package directory

import (
	"github.com/go-ldap/ldap/v3"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
)

func MapSearchResultToUser(result *ldap.Entry) *models.User {
	accountTypeRaw := result.GetAttributeValue("sAMAccountType")

	return &models.User{
		Id:        result.GetAttributeValue("objectGUID"),
		Location:  result.GetAttributeValue("distinguishedName"),
		Upn:       result.GetAttributeValue("userPrincipalName"),
		Email:     result.GetAttributeValue("mail"),
		Name:      result.GetAttributeValue("sAMAccountName"),
		GivenName: result.GetAttributeValue("givenName"),
		Surname:   result.GetAttributeValue("sn"),
		Manager:   result.GetAttributeValue("manager"),
		Type:      MapAccountTypeToDescription(accountTypeRaw),
	}
}

func MapAccountTypeToDescription(accountType string) string {
	knownTypes := map[string]string{
		"268435456":  "SAM_GROUP_OBJECT",
		"268435457":  "SAM_NON_SECURITY_GROUP_OBJECT",
		"536870912":  "SAM_ALIAS_OBJECT",
		"536870913":  "SAM_NON_SECURITY_ALIAS_OBJECT",
		"805306368":  "SAM_NORMAL_USER_ACCOUNT",
		"805306369":  "SAM_MACHINE_ACCOUNT",
		"805306370":  "SAM_TRUST_ACCOUNT",
		"1073741824": "SAM_APP_BASIC_GROUP",
		"1073741825": "SAM_APP_QUERY_GROUP",
		"2147483647": "SAM_ACCOUNT_TYPE_MAX",
	}

	description, exists := knownTypes[accountType]
	if exists {
		return description
	}
	return accountType
}

func MapSearchResultToGroup(result *ldap.Entry) *models.Group {
	return &models.Group{
		Id:       result.GetAttributeValue("objectGUID"),
		Location: result.GetAttributeValue("distinguishedName"),
		Name:     result.GetAttributeValue("sAMAccountName"),
		Type:     result.GetAttributeValue("groupType"),
	}
}
