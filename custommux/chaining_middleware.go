package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/justinas/alice"
)

func main() {
	var (
		PORT             = ":8000"
		mainLogicHandler = http.HandlerFunc(mainLogic)
		chain            = alice.New(filterContentType, setServerTimeCookie).Then(mainLogicHandler)
	)

	http.Handle("/cities", chain)
	fmt.Printf("Server listening at http://localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
