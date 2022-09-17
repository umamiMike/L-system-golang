package main

import (
	"fmt"
	"strings"
)

type System struct {
	axiom      string //the start, must be one of the keys in the rules
	iterations int    //the number of times to recurse through the set
	outfile    string
	rules      string
}

func (s System) generate() strings.Builder {

	// every time we iterate
	// we run through each character of the string
	var bs strings.Builder
	bs.WriteString(s.axiom)
	rs, _ := get_ruleset(s.rules)
	ruleslice := []string{}
	for k, v := range rs.rules {
		ruleslice = append(ruleslice, k, v)
	}
	repl := strings.NewReplacer(ruleslice...)
	fmt.Println("repl: ", repl)

	for n := 0; n <= s.iterations; n++ {
		fmt.Println(repl.Replace(bs.String()))
		bs.WriteString(repl.Replace(bs.String()))
	}
	return bs
}

func get_ruleset(rs string) (*ruleset, error) {
	return select_ruleset(rs)
}
