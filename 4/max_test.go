package max_test

import (
	"max"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	f := func(a, b int) bool {
		result := max.Max(a, b)

		if result < a || result < b {
			return false
		}

		if result != a && result != b {
			return false
		}

		return true
	}

	err := quick.Check(f, nil)
	assert.NoError(t, err)
}
