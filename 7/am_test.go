package am_test

import (
	"am"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type FakeClock struct {
	fixed time.Time
}

func (f FakeClock) Now() time.Time {
	return f.fixed
}

func TestIsMorning(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
		want bool
	}{
		{"0:00", time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), true},
		{"8:00", time.Date(2025, 1, 1, 8, 0, 0, 0, time.UTC), true},
		{"11:59", time.Date(2025, 1, 1, 11, 59, 0, 0, time.UTC), true},
		{"12:00", time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC), false},
		{"20:00", time.Date(2025, 1, 1, 20, 0, 0, 0, time.UTC), false},
		{"23:59", time.Date(2025, 1, 1, 23, 59, 0, 0, time.UTC), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := am.IsAM(FakeClock{tt.time})
			assert.Equal(t, tt.want, got)
		})
	}

}
