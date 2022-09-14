package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)


func main() {
	rules := os.Args[1]
	fmt.Println(rules)
	axiom := os.Args[2]

	iters, _ := strconv.Atoi(os.Args[3])
	fmt.Println("iters are ", iters)

	start := time.Now()
	fmt.Println("starting time")
	sys := System{
		iterations: iters,
		rules:      rules,
		axiom:      axiom,
	}
	fmt.Println("sys", sys)


	grammer := sys.generate()
	fmt.Println("\n\n", grammer)
	fmt.Println(time.Since(start))
	drawSpike("sooky.png", 500, 500, "")
}
