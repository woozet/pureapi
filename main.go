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
		fmt.Println("Header :", r.Header)
		fmt.Println("URL :", r.URL)
		fmt.Println("Path :", r.URL.Path)
		fmt.Println("RawQuery :", r.URL.RawQuery)

		fmt.Fprintf(w, "{ \"path\": \"%s\" }", r.URL.Path[1:])	
	}
    
}

func main() {
    http.HandleFunc("/", handler)
    log.Printf("API Server Ready")
    log.Fatal(http.ListenAndServe(":8080", nil))
}