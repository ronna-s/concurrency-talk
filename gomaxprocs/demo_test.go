package gomaxprocs_test

import (
	"testing"

	"github.com/ronna-s/concurrency-talk/gomaxprocs"
)

func BenchmarkAtomicStorm(b *testing.B) {
	gomaxprocs.AtomicStorm(b.N, 1000)
}

func BenchmarkMutexStorm(b *testing.B) {
	gomaxprocs.MutexStorm(b.N, 1000)
}
