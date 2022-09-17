package main

import (
	"fmt"
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
	fmt.Println(output.String())
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
	fmt.Println(output.String())
	assert.IsType(t, output, bb)

}
