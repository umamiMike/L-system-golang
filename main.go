/*
http://www.cs.unm.edu/~joel/PaperFoldingFractal/L-system-rules.html

An L-system is a formal grammar consisting of 4 parts:

A set of variables: symbols that can be replaced by production rules (see below). In the Fractal Grower software, variables can be any of the 26 lower case English letters a through z

A set of constants: symbols that do not get replaced. In the Fractal Grower software, the constants are any of the following symbols: !, [, ], +, -.

A single axiom which is a string composed of some number of variables and/or constants. The axiom is the initial state of the system.

A set of production rules defining the way variables can be replaced with combinations of constants and other variables. A production consists of two strings - the predecessor and the successor.

*/

package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/spf13/cobra"
	"os"
)

var iters int
var outfile string
var fullgrammer string

type system struct {
	vars       map[string]string
	constants  []string
	axiom      string //the start, must be one of the keys in the rules
	iterations int
	drawing    *gg.Context
	outfile    string
}

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
		str := string(r)

		if contains(constantset, str) {
			newstr = str
		} else {
			newstr = rules[string(r)]
		}
	}

	return newstr
}

func (s system) run() {

	start := s.axiom

	var alliterations []string
	alliterations = append(alliterations, start)

	//for every iteration
	//j
	for n := 0; n <= s.iterations; n++ {

		substring := alliterations[n]

		newsubstring := iterate(substring)
		substring = substring + newsubstring
		alliterations = append(alliterations, substring)
	}

	fmt.Println(alliterations)

	s.drawing.SavePNG(s.outfile)
}

func runsystem(dc *gg.Context) *cobra.Command {

	return &cobra.Command{

		Use: "run",
		RunE: func(cmd *cobra.Command, args []string) error {

			sys := system{
				axiom:      "a",
				iterations: iters,
				vars:       rules,
				constants:  constantset,
				drawing:    dc,
				outfile:    outfile,
			}

			sys.run()

			return nil
		},
	}
}
func main() {
	const W = 1024
	const H = 1024

	dc := gg.NewContext(W, H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	cmd := &cobra.Command{
		Use:          "lsys",
		Short:        "Lsystem grammer generation",
		SilenceUsage: true,
	}
	runsys := runsystem(dc)
	runsys.Flags().IntVarP(&iters, "iterations", "i", 1, "number of iterations")
	runsys.Flags().StringVarP(&outfile, "outfile", "o", "snart.png", "png file to write to")
	cmd.AddCommand(runsys)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
