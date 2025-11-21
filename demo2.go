package demo

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func SpecialCb(ctx context.Context) {
	var muA, muB sync.Mutex

	done := make(chan struct{})
	//fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	muA.Lock()
	// Spawn a goroutine that locks B then A (opposite order).
	go func() {
		defer close(done)
		muB.Lock()
		defer muB.Unlock()

		muA.Lock()
		muA.Unlock()
		//fmt.Println("goroutine: acquired B then A")
	}()

	var x int = 1
	for i := range rand.Intn(1000) {
		x *= i
	}
	_ = x
	// With GOMAXPROCS=1, this line usually runs before the goroutine gets CPU,
	// so main acquires B as well and there's no deadlock.
	// With GOMAXPROCS>=2, the goroutine can grab B first, and we deadlock:
	//   - main holds A, waiting for B
	//   - goroutine holds B, waiting for A

	// Main goroutine locks A first.

	muB.Lock()

	// If we got here, we avoided deadlock (likely GOMAXPROCS=1).
	//fmt.Println("main: acquired A then B")

	// Clean up and let the goroutine finish.

	muB.Unlock()
	muA.Unlock()
	<-done
	//fmt.Println("done")
}
func DoConcurrently(nGoRoutines int, cb func(ctx context.Context)) {
	var wg sync.WaitGroup
	wg.Add(nGoRoutines)
	var i uint32
	go func() {
		for {
			time.Sleep(time.Second)
			if true {
				fmt.Println("> completed:", i)
			}
		}
	}()
	for range nGoRoutines {
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			defer wg.Done()
			defer func() {
				atomic.AddUint32(&i, 1)
			}()
			cb(ctx)
		}()
	}

	wg.Wait()
}
