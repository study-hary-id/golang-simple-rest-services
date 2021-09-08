package main

import (
	"fmt"
	"log"
	"net/http"
)

// CustomServeMux is a multiplexer that handle many HTTP request.
type CustomServeMux struct{}

// ServeHTTP handle HTTP request and serve HTTP response for CustomServeMux.
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
		PORT = ":8000"
	)
	fmt.Printf("Server listening at http://localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
