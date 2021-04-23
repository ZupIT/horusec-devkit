package utils

import (
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/utils/validation/enums"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
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
