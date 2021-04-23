package utils

import (
	"regexp"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/ZupIT/horusec-devkit/pkg/utils/validation/enums"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
)

func CheckInvalidLdapGroups(authType auth.AuthenticationType, groups, permissions []string) error {
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

func PasswordValidationRules() []validation.Rule {
	return []validation.Rule{
		validation.Required,
		validation.Length(8, 255),
		validation.Match(regexp.MustCompile(`[A-Z]`)).Error(enums.MessageMustContainUppercaseCharacter),
		validation.Match(regexp.MustCompile(`[a-z]`)).Error(enums.MessageMustContainLowercaseCharacter),
		validation.Match(regexp.MustCompile(`[0-9]`)).Error(enums.MessageMustContainNumericCharacter),
		validation.Match(regexp.MustCompile(`[!@#$&*-._]`)).Error(enums.MessageMustContainEspecialCharacter),
	}
}
