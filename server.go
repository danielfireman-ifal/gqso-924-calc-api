package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/soma/", somaHandler)
	http.HandleFunc("/sub/", subHandler)
	http.HandleFunc("/mult/", multHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
