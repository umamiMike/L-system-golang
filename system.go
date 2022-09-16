package main

import (
	"bytes"
	"fmt"
	"strings"
)

type System struct {
	axiom      string //the start, must be one of the keys in the rules
	iterations int    //the number of times to recurse through the set
	outfile    string
	rules      string
}

func (s System) generate() bytes.Buffer {

	// every time we iterate
	// we run through each character of the string
	var bs bytes.Buffer
	bs.WriteString(s.axiom)
	ruls, err := get_ruleset(s.rules)
	for n := 0; n <= s.iterations; n++ {
		for _, v := range bs.String() {

			if err != nil {
				fmt.Println("failed to get ruleset")
			}
			//if the rules map has a key then
			// write the value to the buffer
			// if ruls.rules[string(v)] == "" {
			// 	bs.WriteString(ruls.rules[string(v)])
			// }

			if val, ok := ruls.rules[string(v)]; ok {
				bs.WriteString(val)

			}

			if strings.Contains(ruls.constants, string(v)) {
				bs.WriteString(string(v))
			}
		}

	}

	return bs
}

func get_ruleset(rs string) (*ruleset, error) {
	return select_ruleset(rs)
}
