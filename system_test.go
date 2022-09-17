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
	assert.Equal(t, valid, output)

}

func TestCantor(t *testing.T) {
	sys := System{
		iterations: 2,
		rules:      "cantor",
		axiom:      "a",
	}
	output := sys.generate()
	valid := `ababbbaba`
	assert.Equal(t, valid, output)

}

func TestBinaryTree(t *testing.T) {
	sys := System{
		iterations: 2,
		rules:      "binary_tree",
		axiom:      "0",
	}
	output := sys.generate()
	valid := `11[1[0]0]1[0]0`
	// assert.EqualValues(t, ctv, output)
	assert.Equal(t, valid, output)

}
