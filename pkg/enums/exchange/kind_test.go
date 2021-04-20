package exchange

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKind_ToString(t *testing.T) {
	t.Run("Should parse kind exchange with success", func(t *testing.T) {
		assert.Equal(t, "topic", Topic.ToString())
		assert.Equal(t, "fanout", Fanout.ToString())
	})
}