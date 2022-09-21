package directory

import (
	"github.com/go-ldap/ldap/v3"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
	"strconv"
	"time"
)

func getUserAttributes() []string {
	return []string{
		"objectGUID",
		"sAMAccountName",
		"mail",
		"userPrincipalName",
		"givenName",
		"sn",
		"distinguishedName",
		"manager",
		"sAMAccountType",
		"company",
		"department",
		"whenCreated",
		"whenChanged",
		"logonCount",
		"badPwdCount",
		"badPasswordTime",
		"pwdLastSet",
		"lastLogon",
		"lastLogonTimestamp",
	}
}
func getGroupAttributes() []string {
	return []string{
		"objectGUID",
		"sAMAccountName",
		"groupType",
		"distinguishedName",
	}
}

func MapSearchResultToUser(result *ldap.Entry) *models.User {
	accountTypeRaw := result.GetAttributeValue("sAMAccountType")

	return &models.User{
		Id:            result.GetAttributeValue("objectGUID"),
		Location:      result.GetAttributeValue("distinguishedName"),
		Upn:           result.GetAttributeValue("userPrincipalName"),
		Email:         result.GetAttributeValue("mail"),
		Name:          result.GetAttributeValue("sAMAccountName"),
		GivenName:     result.GetAttributeValue("givenName"),
		Surname:       result.GetAttributeValue("sn"),
		Manager:       result.GetAttributeValue("manager"),
		Type:          MapAccountTypeToDescription(accountTypeRaw),
		Company:       result.GetAttributeValue("company"),
		Department:    result.GetAttributeValue("department"),
		CreatedAt:     ParseLdapDate(result.GetAttributeValue("whenCreated")),
		LastUpdatedAt: ParseLdapDate(result.GetAttributeValue("whenChanged")),
		CredentialInfo: &models.UserCredentialInfo{
			FailedLoginAttempts:    ParseIntValue(result.GetAttributeValue("badPwdCount")),
			LastFailedLoginAttempt: ParseLdapDate(result.GetAttributeValue("badPasswordTime")),
			LastLogin:              getLastLogin(result),
			LoginCount:             ParseIntValue(result.GetAttributeValue("logonCount")),
			PasswordLastSet:        ParseLdapDate(result.GetAttributeValue("pwdLastSet")),
		},
	}

}

func getLastLogin(entry *ldap.Entry) time.Time {
	lastLogon := ParseLdapDate(entry.GetAttributeValue("lastLogon"))
	lastLogonTimestamp := ParseLdapDate(entry.GetAttributeValue("lastLogonTimestamp"))

	if lastLogon.After(lastLogonTimestamp) {
		return lastLogon
	}

	return lastLogonTimestamp
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

func ParseLdapDate(value string) time.Time {
	date, _ := time.Parse("20060102150405.0Z07", value)

	return date
}

func ParseIntValue(value string) int {
	parsed, _ := strconv.Atoi(value)

	return parsed
}
