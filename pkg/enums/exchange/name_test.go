package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName_ToString(t *testing.T) {
	t.Run("should parse exchange name with success", func(t *testing.T) {
		assert.Equal(t, "new-analysis", NewAnalysis.ToString())
	})
}
