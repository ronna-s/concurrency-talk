package demo_test

import (
	"testing"
	"testing/synctest"

	"github.com/ronna-s/concurrency-talk"
)

func BenchmarkAtomicStorm(b *testing.B) {
	demo.AtomicStorm(b.N, 1000)
}

func BenchmarkMutexStorm(b *testing.B) {
	demo.MutexStorm(b.N, 1000)
}

func TestDoConcurrently(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		demo.DoConcurrently(1000000, demo.SpecialCb)
		// Wait will block until the goroutine above has finished.
		synctest.Wait()
	})
}
