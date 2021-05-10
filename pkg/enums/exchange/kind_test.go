package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKind_ToString(t *testing.T) {
	t.Run("should parse exchange kind with success", func(t *testing.T) {
		assert.Equal(t, "topic", Topic.ToString())
		assert.Equal(t, "fanout", Fanout.ToString())
	})
}
