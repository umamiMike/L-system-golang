package main

import "fmt"


func main() {

	sys := System{
		iterations: 10,
		rules:      "algea",
		axiom:      "a",
	}
	grammer := sys.generate()
	fmt.Print(grammer)
}
