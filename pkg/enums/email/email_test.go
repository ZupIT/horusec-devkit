package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	t.Run("should return 4 valid values", func(t *testing.T) {
		assert.Len(t, Values(), 4)
	})
}

func TestToString(t *testing.T) {
	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "account-confirmation", AccountConfirmation.ToString())
		assert.Equal(t, "reset-password", ResetPassword.ToString())
		assert.Equal(t, "organization-invite", OrganizationInvite.ToString())
		assert.Equal(t, "repository-invite", RepositoryInvite.ToString())
	})
}
