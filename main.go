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
	constants  map[string]struct{}
	axiom      string //the start, must be one of the keys in the rules
	iterations int
	drawing    *gg.Context
	outfile    string
}

func (s system) run() {
	substring := s.axiom
	for n := 0; n <= s.iterations; n++ {
		for _, r := range substring {
			_, ok := s.constants[string(r)] //checking constants map for existence
			if ok {
				substring += string(r)
			} else {
				substring += rules[string(r)]
			}
			fmt.Println(substring)
		}
	}
	s.drawing.SavePNG(s.outfile)
}

func runsystem(dc *gg.Context) *cobra.Command {

	return &cobra.Command{

		Use: "runsys",
		RunE: func(cmd *cobra.Command, args []string) error {

			sys_to_run := system{
				axiom:      "a",
				iterations: iters,
				vars:       rules,
				constants:  constantset,
				drawing:    dc,
				outfile:    outfile,
			}
			sys_to_run.run()

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
