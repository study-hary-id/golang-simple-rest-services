package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// CustomServeMux is a multiplexer that handles many HTTP requests.
type CustomServeMux struct {
	Started time.Time
}

// ServeHTTP handles HTTP request and serves HTTP response for CustomServeMux.
func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	currentUptime := fmt.Sprintf("Current uptime: %s\n", time.Since(p.Started))
	_, err := w.Write([]byte("CustomServeMux\n" + currentUptime))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return

}

// NewUptimeHandler returns CustomServeMux with provided started time.
func NewUptimeHandler() http.Handler {
	return &CustomServeMux{Started: time.Now()}
}

func main() {
	port := ":8000"
	http.Handle("/", NewUptimeHandler())
	fmt.Printf("Server listening at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
