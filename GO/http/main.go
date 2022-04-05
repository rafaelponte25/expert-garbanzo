package main

import (
	"log"
	"net/http"
)

// Req: http://localhost:1234/upper?word=abc
// Res: ABC
func upperCaseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func main() {
	http.HandleFunc("/upper", upperCaseHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
