package main

import (
	"fmt"
	"os"
)

type System struct {
	axiom      string //the start, must be one of the keys in the rules
	iterations int    //the number of times to recurse through the set
	outfile    string
	rules      string
}

func (s System) generate() []string {

	start := s.axiom
	fmt.Println("start ", start )


	var alliterations []string

	alliterations = append(alliterations, start)
fmt.Println("alliterations :" , alliterations )


	rs, err := get_ruleset(s.rules)


	if err != nil {
		os.Exit(1)
	}

	fmt.Println("s.iterations : " , s.iterations )


	for n := 0; n <= s.iterations; n++ {
		substring := alliterations[n]
		newsubstring := substring + iterate(substring, rs)

		alliterations = append(alliterations, newsubstring)
	}
	fmt.Print("the rules are:\t", s.rules, "\nand the system is:\t ", s, "\nthe ruleset is:\t", rs)

	return alliterations
}

/*
 * utility if a string slice contains the string return true
 */
func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func iterate(s string, ruleset *ruleset) string {
	var newstr string
	fmt.Println("s : " , s )

	for _, r := range s { //for each character in the string
		str := string(r)

		if contains(ruleset.constants, str) {
			newstr = str
		} else {
			newstr = ruleset.rules[string(r)]
		}
	}
	return newstr
}

func get_ruleset(rs string) (*ruleset, error) {
	return select_ruleset(rs)
}
