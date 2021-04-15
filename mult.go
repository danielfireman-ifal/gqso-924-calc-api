package main

import (
	"net/http"
)

func mult(a, b float64) float64 {
	return a * b
}

func multHandler(
	w http.ResponseWriter,
	r *http.Request) {

}
