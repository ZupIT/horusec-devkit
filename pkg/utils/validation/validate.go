// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
