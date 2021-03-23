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
