package main

import (
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(3., soma(1, 2))
}

func TestSomaHandler_Ok(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test.com/soma/1/2",
		nil)
	w := httptest.NewRecorder()
	somaHandler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	result, _ := strconv.ParseFloat(string(body), 64)
	assert.Equal(t, 3., result)
}

func TestSomaHandler_Error2Param(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test.com/soma/1/4d",
		nil)
	w := httptest.NewRecorder()
	somaHandler(w, req)
	resp := w.Result()
	assert.Equal(t, 400, resp.StatusCode)
}

func TestSomaHandler_Error1Param(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test.com/soma/4d/1",
		nil)
	w := httptest.NewRecorder()
	somaHandler(w, req)
	resp := w.Result()
	assert.Equal(t, 400, resp.StatusCode)
}
