package invertmap_test

import (
	"invertmap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertMap(t *testing.T) {
	tests := []struct {
		name string
		in   map[string]int
		want map[int]string
	}{
		{"one", map[string]int{"a": 1}, map[int]string{1: "a"}},
		{"multiple", map[string]int{"a": 1, "b": 2}, map[int]string{1: "a", 2: "b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := invertmap.InvertMap(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}

}
