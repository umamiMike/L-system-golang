package main

var rules = map[string]string{
	"a": "a_b",
	"b": "bc_ab",
	"c": "ca_b",
}
var constantset = map[string]struct{}{
	" ": struct{}{},
	"_": struct{}{},
	"!": struct{}{},
	"<": struct{}{},
	".": struct{}{},
}
