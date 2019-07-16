package main

import "fmt"

func (s system) run() {

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
		fmt.Print(v + " ")
	}
}
