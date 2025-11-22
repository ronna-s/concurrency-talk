package demo

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

const (
	// Max body size we’ll accept (64 MB here – adjust as needed).
	maxBodySize = 64 << 20 // 64 * 1024 * 1024
)

// Response describes what we send back.
type Response struct {
	Array [100]int
}
type Request struct {
	Slice []int
}

func StartService(addr string) (string, func() error, func()) {
	mux := http.NewServeMux()
	mux.HandleFunc("/analyze", AnalyzeHandler)

	server := &http.Server{
		Handler: mux,
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		if err := server.Serve(ln); err != nil {
			//log.Println(err)
		}
	}()
	return ln.Addr().String(), server.Close, func() { <-done }
}

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var data Request
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var resp Response
	writeJSON(w, http.StatusOK, resp)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}
