package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMult(t *testing.T) {
	assert.Equal(t, 4., mult(2, 2))
}
