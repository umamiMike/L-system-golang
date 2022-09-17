package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawOutputsFile(t *testing.T) {
	fname := "testOutputFile.png"
	outfile := drawImage(fname, 500, 500, "ab[]a")
	assert.Equal(t, fname, outfile, "spike of comparison")
}
