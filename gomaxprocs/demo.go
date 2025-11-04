package gomaxprocs

import (
	"sync"
	"sync/atomic"
)

func AtomicStorm(workers, iters int) {
	var counter int64
	var wg sync.WaitGroup
	wg.Add(workers)
	for range workers {
		go func() {
			for range iters {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	_ = counter
}

func MutexStorm(workers, iters int) {
	var (
		x  int64
		mu sync.Mutex
		wg sync.WaitGroup
	)
	wg.Add(workers)
	for range workers {
		go func() {
			for range iters {
				mu.Lock()
				x++
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	_ = x
}
