package analisis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	t.Run("should return 3 valid values", func(t *testing.T) {
		assert.Len(t, Values(), 3)
	})
}
