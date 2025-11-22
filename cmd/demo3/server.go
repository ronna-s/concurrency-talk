package main

import (
	"fmt"
	"time"

	demo "github.com/ronna-s/concurrency-talk"
)

func main() {
	addr, shutdown, wait := demo.StartService("127.0.0.1:8080")
	fmt.Println("listening on:", addr)
	time.Sleep(10 * time.Second)
	_ = shutdown()
	wait()
}
