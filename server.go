package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/soma/", somaHandler)
	http.HandleFunc("/sub/", subHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
