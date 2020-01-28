package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawOutputsFile(t *testing.T) {
	fname := "testOutputFile.png"
	outfile := drawSpike(fname)
	assert.Equal(t, fname, outfile, "spike of comparison")
	os.Open(fname)

}
