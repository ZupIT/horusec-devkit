package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToStringIsAuthorizedType(t *testing.T) {
	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "IsApplicationAdmin", ApplicationAdmin.ToString())
	})

	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "IsCompanyAdmin", CompanyAdmin.ToString())
	})

	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "IsCompanyMember", CompanyMember.ToString())
	})

	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "IsRepositoryAdmin", RepositoryAdmin.ToString())
	})

	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "IsRepositorySupervisor", RepositorySupervisor.ToString())
	})

	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "IsRepositoryMember", RepositoryMember.ToString())
	})
}

func TestValuesIsAuthorizedType(t *testing.T) {
	t.Run("should return 6 valid values", func(t *testing.T) {
		assert.Len(t, Values(), 6)
	})
}
