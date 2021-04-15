package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func mult(a, b float64) float64 {
	return a * b
}

func multHandler(
	w http.ResponseWriter,
	r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	op1, err := strconv.ParseFloat(path[2], 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	op2, err := strconv.ParseFloat(path[3], 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	result := mult(op1, op2)
	fmt.Fprintf(w, "%f", result)
}
