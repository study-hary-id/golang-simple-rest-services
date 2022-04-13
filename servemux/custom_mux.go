package main

import (
	"fmt"
	"log"
	"net/http"
)

// CustomServeMux is a multiplexer that handles many HTTP requests.
type CustomServeMux struct{}

// ServeHTTP handles HTTP request and serves HTTP response for CustomServeMux.
func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		_, err := w.Write([]byte("CustomServeMux\n"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.NotFound(w, r)
	return
}

func main() {
	var (
		mux  = &CustomServeMux{}
		port = ":8000"
	)
	fmt.Printf("Server listening at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
