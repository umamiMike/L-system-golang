package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterate(t *testing.T) {

	sys := System{
		iterations: 10,
		rules:      "algea",
		axiom:      "a",
	}
	assert.True(t, sys.generate())

}
