package utils

import (
	"strings"

	"github.com/ZupIT/horusec-devkit/pkg/utils/validation/enums"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
)

func CheckInvalidLdapGroups(authType auth.AuthorizationType, groups, permissions []string) error {
	if authType == auth.Ldap && isInvalidGroup(groups, permissions) {
		return enums.ErrorInvalidLdapGroup
	}

	return nil
}

func isInvalidGroup(groups, permissions []string) bool {
	for _, group := range groups {
		if group == "" {
			continue
		}

		for _, permission := range permissions {
			if strings.TrimSpace(group) == permission {
				return false
			}
		}
	}

	return true
}
