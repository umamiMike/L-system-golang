package main

var rules = map[string]string{
	"a": "b",
	"b": "c_ab",
	"c": "c!da_b",
	"d": "ca!_b",
}
var constantset = map[string]struct{}{
	" ": struct{}{},
	"_": struct{}{},
	"!": struct{}{},
	"<": struct{}{},
	".": struct{}{},
}
