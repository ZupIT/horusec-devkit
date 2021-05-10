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
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/utils/validation/enums"
)

func TestIsInvalidLdapGroups(t *testing.T) {
	t.Run("should return false when valid group and ldap auth", func(t *testing.T) {
		assert.NoError(t, CheckInvalidLdapGroups(auth.Ldap, []string{"test"}, []string{"test"}))
	})

	t.Run("should return true when invalid group and ldap auth", func(t *testing.T) {
		assert.Error(t, CheckInvalidLdapGroups(auth.Ldap, []string{"test"}, []string{"test2"}))
	})

	t.Run("should return true when invalid group and ignore empty", func(t *testing.T) {
		assert.Error(t, CheckInvalidLdapGroups(auth.Ldap, []string{""}, []string{"", "test"}))
	})

	t.Run("should return false when invalid group and ignore empty", func(t *testing.T) {
		assert.NoError(t, CheckInvalidLdapGroups(auth.Ldap, []string{"", "test"}, []string{"", "test"}))
	})

	t.Run("should return false when not ldap auth", func(t *testing.T) {
		assert.NoError(t, CheckInvalidLdapGroups(auth.Horusec, []string{"test"}, []string{"test"}))
	})
}

func TestPasswordValidationRules(t *testing.T) {
	t.Run("should return no error when valid password", func(t *testing.T) {
		assert.NoError(t, validation.Validate("$3cur3Pa$$?", PasswordValidationRules()...))
	})

	t.Run("should return error when missing uppercase character", func(t *testing.T) {
		err := validation.Validate("insecure!123", PasswordValidationRules()...)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), enums.MessageMustContainUppercaseCharacter)
	})

	t.Run("should return error when missing lowercase character", func(t *testing.T) {
		err := validation.Validate("INSECURE!123", PasswordValidationRules()...)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), enums.MessageMustContainLowercaseCharacter)
	})

	t.Run("should return error when missing numeric character", func(t *testing.T) {
		err := validation.Validate("Insecure!", PasswordValidationRules()...)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), enums.MessageMustContainNumericCharacter)
	})

	t.Run("should return error when especial numeric character", func(t *testing.T) {
		err := validation.Validate("InsecurePass123", PasswordValidationRules()...)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), enums.MessageMustContainEspecialCharacter)
	})
}
