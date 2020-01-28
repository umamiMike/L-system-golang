package main

import "errors"

type ruleset struct {
	rules     map[string]string
	constants []string
}

func select_ruleset(rs_name string) (*ruleset, error) {
	switch rs_name {
	case "algea":
		return algea_set(), nil
	case "koch_curve":
		return koch_curve(), nil
	default:
		return nil, errors.New("please select a valid ruleset")
	}
}

func algea_set() *ruleset {
	rs := ruleset{}
	rs.rules = map[string]string{"a": "ab", "b": "ca", "c": "ba"}
	rs.constants = []string{" ", "_", "!", "<", "."}
	return &rs
}
func koch_curve() *ruleset {
	rules := ruleset{}
	rules.rules = map[string]string{"F": "F+F-F-F+F"}
	rules.constants = []string{"+", "-"}
	return &rules
}

// ./notes.txt
func fractal_plant() *ruleset {
	rules := ruleset{}
	rules.rules = map[string]string{"X": "F+[[X]-X]-F[-FX]+X", "F": "FF"}
	rules.constants = []string{"+", "[", "]", "-"}
	return &rules
}
