package main

import (
	"bytes"
	"errors"
)

type ruleset struct {
	rules     map[string]string
	constants string
}

func select_ruleset(rs_name string) (*ruleset, error) {
	switch rs_name {
	case "algea":
		return algea_set(), nil
	case "koch_curve":
		return koch_curve(), nil
	case "plant":
		return fractal_plant(), nil
	case "cantor":
		return cantor_set(), nil
	case "binary_tree":
		return binary_tree(), nil
	case "test":
		return test(), nil
	default:
		return nil, errors.New("please select a valid ruleset")
	}
}

func algea_set() *ruleset {
	rs := ruleset{}
	rs.rules = map[string]string{"a": "ab", "b": "a"}
	rs.constants = "_!<."
	return &rs
}
func koch_curve() *ruleset {
	rules := ruleset{}
	rules.rules = map[string]string{"F": "F+F-F-F+F"}
	rules.constants = "+-"
	return &rules
}

func cantor_set() *ruleset {
	rules := ruleset{}
	rules.rules = map[string]string{"a": "aba", "b": "bbb"}
	rules.constants = ""
	return &rules
}
func fractal_plant() *ruleset {
	rules := ruleset{}
	rules.rules = map[string]string{"X": "F+[[X]-X]-F[-FX]+X",
		"F": "FF",
	}
	rules.constants = "+[]-"
	return &rules
}

func binary_tree() *ruleset {
	rules := ruleset{}
	rules.rules = map[string]string{"1": "11",
		"0": "1[0]0",
	}
	rules.constants = "[]"
	return &rules
}

func test() *ruleset {
	rules := ruleset{}
	rules.rules = map[string]string{
		"a":   "bat",
		"b":   "bet",
		"n":   "ab",
		"t":   "tan",
		"bet": "rootdown",
	}
	rules.constants = "+[]-"
	return &rules
}

func (rs *ruleset) Contains(s string) bool {
	return false
}

func (rs *ruleset) Rules() string {
	var rb bytes.Buffer
	rb.WriteString("rule characters: ")
	for k, _ := range rs.rules {
		rb.WriteString(k)
		rb.WriteString(", ")
	}
	rb.WriteString("constants: ")
	for _, v := range rs.constants {
		rb.WriteString(string(v))
		rb.WriteString(", ")
	}

	return rb.String()
}
