package main

import (
	"os"
	"strings"

	"github.com/umamimike/L-system-golang/db"
)

var dbinstance = db.NewBadger("/tmp/badger")

type System struct {
	axiom      string //the start, must be one of the keys in the rules
	iterations int    //the number of times to recurse through the set
	outfile    string
	rules      string
}

func (s System) generate() []string {

	start := s.axiom // one of the keys contained in the ruleste rules map
	var alliterations []string
	alliterations = append(alliterations, start)

	rs, err := get_ruleset(s.rules)
	if err != nil {
		os.Exit(1)
	}

	for n := 0; n <= s.iterations; n++ {
		substring := alliterations[n]
		substring = substring + iterate(substring, rs)
		alliterations = append(alliterations, substring)
	}
	//	drawSpike(s.outfile)
	dbinstance.Write("fookin_data", strings.Join(alliterations, " "))
	//println(string(get("fookin_data")))
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
