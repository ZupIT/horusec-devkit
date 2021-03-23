package confidence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	t.Run("should return 3 valid values", func(t *testing.T) {
		assert.Len(t, Values(), 3)
	})
}

func TestToString(t *testing.T) {
	t.Run("should success parse all values to string", func(t *testing.T) {
		assert.Equal(t, "HIGH", High.ToString())
		assert.Equal(t, "MEDIUM", Medium.ToString())
		assert.Equal(t, "LOW", Low.ToString())
	})
}
