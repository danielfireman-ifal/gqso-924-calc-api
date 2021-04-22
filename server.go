package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/soma/", somaHandler)
	http.HandleFunc("/sub/", subHandler)
	http.HandleFunc("/mult/", multHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}
