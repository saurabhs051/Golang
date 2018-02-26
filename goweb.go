package main

import (
	"fmt"
	"log"
	"net/http"
) //library for networking

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expert web design by Saurabh")
}

func main() {
	http.HandleFunc("/", index_handler) //http handler //index_handler is a func
	http.HandleFunc("/about/", about_handler)
	err := http.ListenAndServe(":8000", nil) //localhost port 8000
	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	}

}
