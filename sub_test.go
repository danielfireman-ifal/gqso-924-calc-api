package main

import (
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub_OK(t *testing.T) {
	assert.Equal(t, 1., sub(2, 1))
}

func TestSubHandler_OK(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test.com/sub/2/1",
		nil)
	w := httptest.NewRecorder()
	somaHandler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	result, _ := strconv.ParseFloat(string(body), 64)
	assert.Equal(t, 1., result)
}
