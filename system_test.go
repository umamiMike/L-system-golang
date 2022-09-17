package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlgea(t *testing.T) {

	sys := System{
		iterations: 7,
		rules:      "algea",
		axiom:      "a",
	}
	var bb strings.Builder
	output := sys.generate()
	assert.IsType(t, output, bb)

}

func TestTest(t *testing.T) {

	sys := System{
		iterations: 2,
		rules:      "test",
		axiom:      "a",
	}
	var bb strings.Builder
	output := sys.generate()
	assert.IsType(t, output, bb)

}

func TestKoch(t *testing.T) {
	sys := System{
		iterations: 9,
		rules:      "koch_curve",
		axiom:      "F",
	}
	var cb strings.Builder
	output := sys.generate()
	assert.IsType(t, output, cb)
	assert.True(t, true)

}
