package main

import "fmt"

/*
returns bool if the carray of strings contains the string supplied by the second arg
*/
func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func iterate(s string) string {
	var newstr string
	for _, r := range s { //for each character in the string
		var algea_set = algea_set()
		str := string(r)

		if contains(algea_set.constants, str) {
			newstr = str
		} else {
			newstr = algea_set.rules[string(r)]
		}
	}
	return newstr
}
func (s system) run() {

	fmt.Println(s.outfile)
	start := s.axiom
	var alliterations []string
	alliterations = append(alliterations, start)

	for n := 0; n <= s.iterations; n++ {

		substring := alliterations[n]
		newsubstring := iterate(substring)
		substring = substring + newsubstring
		alliterations = append(alliterations, substring)
	}

	for _, v := range alliterations {
		fmt.Println(v + " ")
	}
}
