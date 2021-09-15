package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var (
		PORT             = ":8000"
		mainLogicHandler = http.HandlerFunc(mainLogic)
	)
	http.Handle("/cities", filterContentType(setServerTimeCookie(mainLogicHandler)))
	fmt.Printf("Server listening at http://localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
