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

// This isn't a real test, just using it to step through a real mutex in the debugger
func TestExploreRealMutex(t *testing.T) {
	d := 0
	n_tests := 2
	mu := sync.Mutex{} // Returns a nil struct {state: 0, sema: 0}
	wg := sync.WaitGroup{}
	wg.Add(n_tests)
	for i := 0; i < n_tests; i++ {
		go func(n int) {
			defer wg.Done()
			t.Log(n)
			// under the hood, this is atomically checking if the mutex state = 0 (unlocked) and if so, setting it to a locked state of 1.
			mu.Lock()
			d += 1
			// cannot be called on an already unlocked mutex. also any goroutine can unlock, doesn't have to be the original locker!
			// between the call to lock and unlock, something seems to have modified the mu.State. It increased from 1 after locking to 8 (!).
			mu.Unlock()
		}(i)
	}
	wg.Wait()
}

// Test runtime error calling unlock on an unlocked lock.
// func TestUnlockUnlockedLock(t *testing.T) {
// 	mu := sync.Mutex{} // Returns a nil struct {state: 0, sema: 0}
// 	mu.Unlock()
// }

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
// BenchmarkFootex-8        4704144               259.9 ns/op
// BenchmarkMutex-8         5907390               206.8 ns/op
// PASS
// ok      ap/mutexes      3.219s
