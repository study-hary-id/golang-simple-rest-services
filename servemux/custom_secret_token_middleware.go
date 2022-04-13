package main

import (
	"fmt"
	"log"
	"net/http"
)

// SecretTokenHandler secures a request with a secret token.
type SecretTokenHandler struct {
	next   http.Handler
	secret string
}

// ServeHTTP makes SecretTokenHandler implement the http.Handler interface.
func (p *SecretTokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("secret_token") == p.secret {
		p.next.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	port := ":8000"
	http.Handle("/", &SecretTokenHandler{
		next:   NewUptimeHandler(),
		secret: "middleware",
	})
	fmt.Printf("Server listening at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
