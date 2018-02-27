package main

import (
	"fmt"
	"log"
	"net/http"
)

func firstpage_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1  style='color:red'>This is a running Go web server.</h1>
		<h2><p>A  new paragraph</p></h2>`)
}

func main() {
	http.HandleFunc("/", firstpage_handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	}
}
