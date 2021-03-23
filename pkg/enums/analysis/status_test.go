package analysis

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
	t.Run("should return 3 valid values", func(t *testing.T) {
		assert.Equal(t, "running", Running.ToString())
		assert.Equal(t, "success", Success.ToString())
		assert.Equal(t, "error", Error.ToString())
	})
}
