package utils

import (
	"regexp"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/utils/validation/enums"
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

// nolint // necessary magic number and valid suspicious regex
func PasswordValidationRules() []validation.Rule {
	return []validation.Rule{
		validation.Required,
		validation.Length(8, 255),
		validation.Match(regexp.MustCompile(enums.RegexUppercaseCharacter)).
			Error(enums.MessageMustContainUppercaseCharacter),
		validation.Match(regexp.MustCompile(enums.RegexLowercaseCharacter)).
			Error(enums.MessageMustContainLowercaseCharacter),
		validation.Match(regexp.MustCompile(enums.RegexNumericCharacter)).
			Error(enums.MessageMustContainNumericCharacter),
		validation.Match(regexp.MustCompile(enums.RegexEspecialCharacter)).
			Error(enums.MessageMustContainEspecialCharacter),
	}
}
