package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDrawOutputsFile(t *testing.T) {
	fname := "testOutputFile.png"
	outfile := drawSpike(fname, 10)
	assert.Equal(t, fname, outfile, "spike of comparison")
	os.Open(fname)

}
