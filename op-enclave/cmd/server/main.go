package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/mdlayher/vsock"
)

// small HTTP proxy that forwards requests to a vsock service
func main() {
	pool := sync.Pool{
		New: func() any {
			conn, err := vsock.Dial(16, 1234, &vsock.Config{})
			if err != nil {
				log.Printf("Error dialing vsock: %v", err)
				return nil
			}
			return conn
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusOK)
			return
		}
		r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
		req, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_ = r.Body.Close()

		conn := pool.Get().(*vsock.Conn)
		if conn == nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer pool.Put(conn)

		_, err = conn.Write(req)
		if err != nil {
			log.Printf("Error writing to vsock: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		dec := json.NewDecoder(conn)
		dec.UseNumber()

		var raw json.RawMessage
		if err := dec.Decode(&raw); err != nil {
			log.Printf("Error decoding response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(raw)
	}

	srv := &http.Server{
		Addr:              ":7333",
		Handler:           http.HandlerFunc(handler),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
