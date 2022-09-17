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

func (s System) generate() string {

	// every time we iterate
	// we run through each character of the string
	var bs strings.Builder
	bs.WriteString(s.axiom)
	rs, _ := get_ruleset(s.rules)
	ruleslice := []string{}
	for k, v := range rs.rules {
		ruleslice = append(ruleslice, k, v)
	}
	fmt.Println("ruleset is: ", ruleslice)
	repl := strings.NewReplacer(ruleslice...)

	for n := 0; n <= s.iterations-1; n++ {
		in := bs.String()
		bs.Reset()
		new := repl.Replace(in)
		bs.WriteString(new)
	}
	return bs.String()
}

func get_ruleset(rs string) (*ruleset, error) {
	return select_ruleset(rs)
}
