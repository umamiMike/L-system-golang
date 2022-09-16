package main

import (
	"fmt"
	"flag"
)


func main() {
	rules := flag.String("rules", "algea", "rules (algea koch_curve plant test)") 
	axiom := flag.String("axiom", "F", "axiom, the seed value")

	iterPtr := flag.Int("iterations", 1, " an int")

	flag.Parse()
	// taking a ruleset
	// returns a string listing all of the rules and constants
	rs, _ := select_ruleset(*rules);
	fmt.Println("rs.rules.Keys: " , rs.Rules())
	sys := System{
		iterations: *iterPtr,
		rules:      *rules,
		axiom:      *axiom,
	}
	grammer := sys.generate()
	fmt.Println(grammer.String())
	// fmt.Println(strings.Join(grammer, ""))
	// drawImage("sookyb.png", 1500, 1500, strings.Join(grammer, ""))
}
