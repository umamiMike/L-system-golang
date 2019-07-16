/*
http://www.cs.unm.edu/~joel/PaperFoldingFractal/L-system-rules.html

An L-system is a formal grammar consisting of 4 parts:

A set of variables: symbols that can be replaced by production rules (see
below). In the Fractal Grower software, variables can be any of the 26 lower
case English letters a through z

A set of constants: symbols that do not get replaced. In the Fractal Grower
software, the constants are any of the following symbols: !, [, ], +, -.

A single axiom which is a string composed of some number of variables and/or
constants. The axiom is the initial state of the system.

A set of production rules defining the way variables can be replaced with
combinations of constants and other variables. A production consists of two
strings - the predecessor and the successor.

*/

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type system struct {
	vars       map[string]string
	axiom      string //the start, must be one of the keys in the rules
	iterations int    //the number of times to recurse through the set
	outfile    string
}

var iters int
var outfile string
var input_axiom string

func runsystem() *cobra.Command {

	return &cobra.Command{

		Use: "run",
		RunE: func(cmd *cobra.Command, args []string) error {
			sys := system{
				axiom:      input_axiom,
				iterations: iters,
				outfile:    outfile,
			}
			sys.run()
			return nil
		},
	}
}

func main() {

	outfile := "test.png"

	RootCmd := &cobra.Command{
		Use:          "lsys",
		Short:        "Lsystem grammer generation",
		SilenceUsage: true,
	}

	runsys := runsystem()
	runsys.Flags().IntVarP(&iters, "iterations", "i", 1, "number of iterations")
	runsys.Flags().StringVarP(&outfile, "outfile", "o", "snart.png", "png file to write to")
	runsys.Flags().StringVarP(&input_axiom, "axiom", "a", "a", "which axiom in the ruleset to begin with")

	RootCmd.AddCommand(runsys)
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
