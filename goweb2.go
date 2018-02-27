package main

import (
	"fmt"
	"log"
	"net/http"
)

func firstpage_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a running Go web server.")
}

func main() {
	http.HandleFunc("/", firstpage_handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	}
}
