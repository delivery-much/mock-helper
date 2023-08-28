package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasArgument(t *testing.T) {
	t.Run("Should return false if the call does not have the argument", func(t *testing.T) {
		c := MockCall{}
		arg := 42

		res := c.HasArgument(arg)
		assert.False(t, res)

		c.Args = []any{
			"argument",
		}
		res = c.HasArgument(arg)
		assert.False(t, res)

		c.Args = []any{
			arg,
		}
		res = c.HasArgument(arg)
		assert.True(t, res)
	})
	t.Run("Should return true if the call has the argument", func(t *testing.T) {
		c := MockCall{}
		arg := 42

		c.Args = []any{
			arg,
		}
		res := c.HasArgument(arg)
		assert.True(t, res)

		c.Args = []any{
			"argument",
			arg,
		}
		res = c.HasArgument(arg)
		assert.True(t, res)
	})
}
