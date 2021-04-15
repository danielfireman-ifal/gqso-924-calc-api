package main

import (
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMult(t *testing.T) {
	assert.Equal(t, 4., mult(2, 2))
}

func TestMultHandler_OK(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test.com/mult/2/1",
		nil)
	w := httptest.NewRecorder()
	multHandler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	result, _ := strconv.ParseFloat(string(body), 64)
	assert.Equal(t, 2., result)
}

func TestMultHandler_Error2Param(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test.com/mult/1/4d",
		nil)
	w := httptest.NewRecorder()
	multHandler(w, req)
	resp := w.Result()
	assert.Equal(t, 400, resp.StatusCode)
}

func TestMultHandler_Error1Param(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test.com/mult/4d/1",
		nil)
	w := httptest.NewRecorder()
	multHandler(w, req)
	resp := w.Result()
	assert.Equal(t, 400, resp.StatusCode)
}
