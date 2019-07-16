package main

import ()

type ruleset struct {
	rules     map[string]string
	constants []string
}

func algea_set() *ruleset {
	rules := ruleset{}
	rules.rules = map[string]string{"a": "ab", "b": "ca", "c": "ba"}
	rules.constants = []string{" ", "_", "!", "<", "."}
	return &rules
}
