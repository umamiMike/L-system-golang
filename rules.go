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
func koch_curve() *ruleset {
	rules := ruleset{}
	rules.rules = map[string]string{"F": "F+F-F-F+F"}
	rules.constants = []string{"+", "-"}
	return &rules
}

//Example 7: Fractal plant
//See also: Barnsley fern
//variables : X F
//constants : + − [ ]
//start : X
//rules : (X → F+[[X]-X]-F[-FX]+X), (F → FF)
//angle : 25°
//Here, F means "draw forward", − means "turn left 25°", and + means "turn
//right 25°". X does not correspond to any drawing action and is used to
//control the evolution of the curve. The square bracket "[" corresponds to
//saving the current values for position and angle, which are restored when the
//corresponding "]" is executed.
func fractal_plant() *ruleset {

	rules := ruleset{}
	rules.rules = map[string]string{"X": "F+[[X]-X]-F[-FX]+X", "F": "FF"}
	rules.constants = []string{"+", "[", "]", "-"}
	return &rules
}
