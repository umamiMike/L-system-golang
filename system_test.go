package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlgea(t *testing.T) {

	sys := System{
		iterations: 1,
		rules:      "algea",
		axiom:      "a",
	}
	output := sys.generate()

	foo := `ab`
	assert.Equal(t, foo, output)

}

func TestKoch(t *testing.T) {
	sys := System{
		iterations: 2,
		rules:      "koch_curve",
		axiom:      "F",
	}
	output := sys.generate()
	valid := `F+F-F-F+F+F+F-F-F+F-F+F-F-F+F-F+F-F-F+F+F+F-F-F+F`
	// assert.EqualValues(t, ctv, output)
	assert.Equal(t, valid, output)

}

// func TestTest(t *testing.T) {

// 	sys := System{
// 		iterations: 2,
// 		rules:      "test",
// 		axiom:      "a",
// 	}
// 	output := sys.generate()
// 	assert.True(t, true)
// }
