package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub_OK(t *testing.T) {
	assert.Equal(t, 1, sub(2, 1))
}
