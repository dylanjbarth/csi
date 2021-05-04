package concurrency

import (
	"reflect"
	"sync"
	"testing"
)

var outer = 1000
var inner = 20
var max = outer * inner

func TestConcurrencyStrategies(t *testing.T) {
	var start uint64
	cases := []counterService{
		// &NoSync{start}, // this fails as expected, TODO is there a nice way to wrap this as an expected failure?
		&AtomicCount{start},
		&SyncMutexCount{sync.Mutex{}, start},
		initChannelCounter(start),
	}
	for _, c := range cases {
		// last := start
		results := make(chan uint64, max)
		// Test monotonically increasing by calling multiple times __within__ a goroutine
		for i := 0; i < outer; i++ {
			go func() {
				var last uint64
				for j := 0; j < inner; j++ {
					new := c.getNext()
					if last > 0 && new < last {
						t.Errorf("Testing %s failed. Values didn't monotonically increase. Received %d after %d within the same goroutine.", reflect.TypeOf(c), new, last)
					}
					last = new
					results <- new
				}
			}()
		}
		// drain channel once tests have run and calc max
		var curr uint64
		for i := 0; i < max; i++ {
			n := <-results
			if n > curr {
				curr = n
			}
		}
		if curr != uint64(max) {
			t.Errorf("Testing %s failed. Expected max value to be %d but got %d.", reflect.TypeOf(c), max, curr)
		}
	}
}

func BenchmarkNoSync(b *testing.B) {
	c := &NoSync{0}
	for n := 0; n < b.N; n++ {
		c.getNext()
	}
}

func BenchmarkAtomicCount(b *testing.B) {
	c := &AtomicCount{0}
	for n := 0; n < b.N; n++ {
		c.getNext()
	}
}

func BenchmarkSyncMutexCount(b *testing.B) {
	c := &SyncMutexCount{sync.Mutex{}, 0}
	for n := 0; n < b.N; n++ {
		c.getNext()
	}
}

func BenchmarkChannelCount(b *testing.B) {
	c := initChannelCounter(0)
	for n := 0; n < b.N; n++ {
		c.getNext()
	}
}

// BenchmarkNoSync-8               875835890                1.325 ns/op
// BenchmarkAtomicCount-8          263140364                4.525 ns/op
// BenchmarkSyncMutexCount-8       73563944                15.12 ns/op
// BenchmarkChannelCount-8          3191479               365.9 ns/op
