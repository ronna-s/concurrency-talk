package demo_test

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"runtime"
	"sync"
	"testing"
	"testing/synctest"

	"github.com/ronna-s/concurrency-talk"
)

func BenchmarkAtomicStorm(b *testing.B) {
	old := runtime.GOMAXPROCS(1)
	defer runtime.GOMAXPROCS(old)
	demo.AtomicStorm(b.N, 1000)

}

func BenchmarkMutexStorm(b *testing.B) {
	old := runtime.GOMAXPROCS(1)
	defer runtime.GOMAXPROCS(old)
	demo.MutexStorm(b.N, 1000)
}

func TestDoConcurrently(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		demo.DoConcurrently(1000000, demo.SpecialCb)
		// Wait will block until the goroutine above has finished.
		synctest.Wait()
	})
}

func BenchmarkStartService(b *testing.B) {
	var wg sync.WaitGroup
	for range b.N {
		wg.Go(func() {
			payload := demo.Request{
				Slice: make([]int, rand.Intn(1000)),
			}
			data, _ := json.Marshal(payload)

			req, err := http.NewRequest("POST", "http://127.0.0.1:8080/demo", bytes.NewBuffer(data))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			demo.AnalyzeHandler(resp, req)
			if err != nil {
				b.Error("Error:", err)
				return
			}
			// Check the status code is what we expect.
			if status := resp.Code; status != http.StatusOK {
				b.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			decoder := json.NewDecoder(resp.Body)
			if err := decoder.Decode(&payload); err != nil {
				b.Error("Error:", err)
			}
		})
	}

	wg.Wait()
}
