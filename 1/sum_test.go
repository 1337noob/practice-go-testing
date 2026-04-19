package sum_test

import (
	"sum"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := []struct {
		num1, num2 int
		want       int
	}{
		{1, 1, 2},
		{-10, 15, 5},
		{-3, -4, -7},
	}

	for _, test := range tests {
		got := sum.Sum(test.num1, test.num2)
		assert.Equal(t, test.want, got)
	}
}
