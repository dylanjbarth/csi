package mutex

import (
	"sync"
	"testing"
)

func TestFootex(t *testing.T) {
	tok := footex{}

	// Test concurrent access to shared var.
	n_test := 1000
	d := uint32(0)
	results := make(chan uint32, n_test)
	for i := 0; i < n_test; i++ {
		// t.Logf("Starting %d\n", i)
		go func(n int) {
			// t.Logf("%d calling lock\n", n)
			tok.Lock()
			d += 1
			results <- d
			// t.Logf("%d calling unlock\n", n)
			tok.Unlock()
			// t.Logf("%d returning\n", n)
		}(i)
	}
	// Ensure that we incremented correctly and have the right max
	var max uint32
	for i := 0; i < n_test; i++ {
		r := <-results
		if r > max {
			max = r
		}
	}
	expected := uint32(n_test)
	if max != expected {
		t.Errorf("Expected max value to equal %d but got %d", expected, max)
	}
}

func BenchmarkFootex(b *testing.B) {
	d := 0
	wg := sync.WaitGroup{}
	tok := footex{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func(n int) {
			defer wg.Done()
			tok.Lock()
			defer tok.Unlock()
			d += 1
		}(i)
	}
	wg.Wait()
}

func BenchmarkMutex(b *testing.B) {
	d := 0
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			d += 1
		}()
	}
	wg.Wait()
}

// $ go test -bench=.
// goos: darwin
// goarch: amd64
// pkg: ap/mutexes
// cpu: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
// BenchmarkFootex-8          10000           5123974 ns/op
// BenchmarkMutex-8         5518824               226.1 ns/op
// PASS
// ok      ap/mutexes      53.055s
// So the sync.Mutex approach is 22,662 times faster :)
// There is probably a much cleaner way to do this in pure go.
