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
	"github.com/spf13/cobra"
	"os"
	"time"
)

type system struct {
	vars       map[string]string
	constants  map[string]struct{}
	axiom      string //the start, must be one of the keys in the rules
	iterations int
}

func (system) run() {
	teststring := s.axiom
	substring := ""
	for n := 1; n <= s.iterations; n++ {
		for _, r := range teststring {
			_, ok := constants[string(r)] //checking constants map for existence
			if ok {
				substring += string(r)
				break //add only the value from the const
			}
			substring += rules[string(r)]
		}
		teststring += substring
	}
	fmt.Print(teststring)
}

func runsystem() *cobra.Command {

	return &cobra.Command{

		Use: "runsys",
		RunE: func(cmd *cobra.Command, args []string) error {
			sys_to_run := system{
				axiom:      "a",
				iterations: 12,
				vars:       rules,
				constants:  constantset,
			}
			sysrun.run(rules)
			return nil
		},
	}
}
func main() {
	cmd := &cobra.Command{
		Use:          "lsys",
		Short:        "Lsystem grammer generation",
		SilenceUsage: true,
	}
	cmd.AddCommand(runsystem())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
