// faster_with_lower_gomaxprocs.go
package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var workers int
	var iters int
	var mode string
	flag.IntVar(&workers, "workers", 400, "number of goroutines")
	flag.IntVar(&iters, "iters", 300_000, "increments per goroutine")
	flag.StringVar(&mode, "mode", "atomic", "atomic|mutex")
	flag.Parse()

	fmt.Printf("GOMAXPROCS=%d workers=%d iters=%d mode=%s\n",
		runtime.GOMAXPROCS(0), workers, iters, mode)

	start := time.Now()
	switch mode {
	case "atomic":
		atomicStorm(workers, iters)
	case "mutex":
		mutexStorm(workers, iters)
	default:
		panic("unknown mode")
	}
	fmt.Println("elapsed:", time.Since(start))
}

func atomicStorm(workers, iters int) {
	var counter int64
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			for i := 0; i < iters; i++ {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	_ = counter
}

func mutexStorm(workers, iters int) {
	var x int64
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			for i := 0; i < iters; i++ {
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

//package main
//
//import (
//	"context"
//	"fmt"
//	"sync"
//	"sync/atomic"
//	"time"
//)
//
//func specialCb(ctx context.Context) {
//	var muA, muB sync.Mutex
//	var wg sync.WaitGroup
//	wg.Add(1)
//
//	//fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
//
//	// Spawn a goroutine that locks B then A (opposite order).
//	go func() {
//		defer wg.Done()
//		muB.Lock()
//		defer muB.Unlock()
//
//		muA.Lock()
//		muA.Unlock()
//
//		//fmt.Println("goroutine: acquired B then A")
//	}()
//
//	// With GOMAXPROCS=1, this line usually runs before the goroutine gets CPU,
//	// so main acquires B as well and there's no deadlock.
//	// With GOMAXPROCS>=2, the goroutine can grab B first, and we deadlock:
//	//   - main holds A, waiting for B
//	//   - goroutine holds B, waiting for A
//
//	// Main goroutine locks A first.
//
//	muA.Lock()
//	muB.Lock()
//
//	// If we got here, we avoided deadlock (likely GOMAXPROCS=1).
//	//fmt.Println("main: acquired A then B")
//
//	// Clean up and let the goroutine finish.
//
//	muB.Unlock()
//	muA.Unlock()
//	wg.Wait()
//	//fmt.Println("done")
//}
//func doConcurrently(nGoRoutines int, cb func(ctx context.Context)) {
//	var wg sync.WaitGroup
//	wg.Add(nGoRoutines)
//	var i uint32
//	go func() {
//		for {
//			time.Sleep(time.Second)
//			fmt.Println("> completed:", i)
//		}
//	}()
//	for range nGoRoutines {
//		go func() {
//			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//			defer cancel()
//			defer wg.Done()
//			defer func() {
//				atomic.AddUint32(&i, 1)
//			}()
//			cb(ctx)
//		}()
//	}
//
//	wg.Wait()
//}
//
//func main() {
//	doConcurrently(1000000, specialCb)
//}
