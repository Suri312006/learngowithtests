package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)

	})

	t.Run("it runs safely in a concurrent environment", func(t *testing.T) {
		wantedCount := 10000
		counter := NewCounter()

		// waits for a collections of goroutines to finish
		// the main goroutine calls Add to set the number of goroutines
		// to wait for/
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				// each of the goroutines calls Done when finished
				wg.Done()
			}()
		}

		// blocks until all the goroutines have finished
		wg.Wait()

		assertCounter(t, counter, wantedCount)

	})

}
func NewCounter() *Counter {
	return &Counter{}
}
func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
