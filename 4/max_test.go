package max_test

import (
	"max"
	"testing"
	"testing/quick"
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

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
