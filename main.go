package main

import (
	"fmt"
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "{ \"code\": %d, \"message\": \"method not allowed\" }", http.StatusMethodNotAllowed)
	} else {
		fmt.Fprintf(w, "{ \"path\": \"%s\" }", r.URL.Path[1:])	
	}
    
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}