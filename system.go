package main

import "strings"

/*
returns bool if the array of strings contains the string supplied by the second arg
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
func (s system) run() {

	start := s.axiom
	var alliterations []string
	alliterations = append(alliterations, start)

	for n := 0; n <= s.iterations; n++ {
		substring := alliterations[n]
		substring = substring + iterate(substring, fractal_plant())
		alliterations = append(alliterations, substring)
	}

	//	drawSpike(s.outfile)
	write(strings.Join(alliterations, " "))

}
