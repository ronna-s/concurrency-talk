package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/ronna-s/concurrency-talk/gomaxprocs"
)

func main() {
	var (
		workers int
		iters   int
		mode    string
	)
	
	flag.IntVar(&workers, "workers", 400, "number of goroutines")
	flag.IntVar(&iters, "iters", 300_000, "increments per goroutine")
	flag.StringVar(&mode, "mode", "atomic", "atomic|mutex")
	flag.Parse()

	fmt.Printf("GOMAXPROCS=%d workers=%d iters=%d mode=%s\n",
		runtime.GOMAXPROCS(0), workers, iters, mode)

	start := time.Now()
	switch mode {
	case "atomic":
		gomaxprocs.AtomicStorm(workers, iters)
	case "mutex":
		gomaxprocs.MutexStorm(workers, iters)
	default:
		panic("unknown mode")
	}
	fmt.Println("elapsed:", time.Since(start))
}
