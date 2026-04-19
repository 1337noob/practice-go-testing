package counter_test

import (
	"counter"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	c := counter.Counter{}
	wg := sync.WaitGroup{}
	goroutinesCount := 1000
	incrementCount := 1000

	wg.Add(goroutinesCount)
	for i := 0; i < goroutinesCount; i++ {
		go func() {
			for j := 0; j < incrementCount; j++ {
				c.Inc()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	want := goroutinesCount * incrementCount
	got := c.Value()
	assert.Equal(t, want, got)
}
