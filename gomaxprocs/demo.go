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
