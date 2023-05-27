package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("single thread", func(t *testing.T) {
		c := NewCounter()

		c.Inc()
		c.Inc()
		c.Inc()

		assertCounter(t, c, 3)
	})

	t.Run("multi threaded safety", func(t *testing.T) {
		want := 1000
		c := NewCounter()

		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < 1000; i++ {
			go func() {
				c.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, c, want)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
