package main

import "fmt"

type System struct {
	axiom      string //the start, must be one of the keys in the rules
	iterations int    //the number of times to recurse through the set
	outfile    string
	rules      string
}

func main() {

	sys := System{
		iterations: 10,
		rules:      "algea",
		axiom:      "a",
	}
	grammer := sys.generate()
	fmt.Print(grammer[3])
}
