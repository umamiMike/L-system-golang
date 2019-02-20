package main

import (
	"github.com/fogleman/gg"
	"math/rand"
)

var rules = map[string]string{
	"a": "ab",
	"b": "ca",
	"c": "ba",
}
var constantset = []string{
	" ",
	"_",
	"!",
	"<",
	".",
}
