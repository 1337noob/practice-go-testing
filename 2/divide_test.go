package divide_test

import (
	"divide"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result, err := divide.Divide(10.8, 2)
		assert.NoError(t, err)
		assert.Equal(t, 5.4, result)
	})

	t.Run("error", func(t *testing.T) {
		result, err := divide.Divide(10.8, 0)
		assert.Error(t, err)
		assert.ErrorIs(t, err, divide.ErrDivisionByZero)
		assert.Equal(t, 0.0, result)
	})
}
