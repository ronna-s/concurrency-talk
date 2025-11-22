package main

import (
	"github.com/ronna-s/concurrency-talk"
)

func main() {
	demo.DoConcurrently(1000000, demo.SpecialCb)
}
