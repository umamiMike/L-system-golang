package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

//var dbinstance = db.NewBadger("./tmp/badger")

func main() {
	rules := os.Args[1]
	axiom := os.Args[2]
	iters, _ := strconv.Atoi(os.Args[3])

	fmt.Print(rules)
	start := time.Now()
	sys := System{
		iterations: iters,
		rules:      rules,
		axiom:      axiom,
	}

	grammer := sys.generate()
	fmt.Println("\n\n", grammer)
	fmt.Println(time.Since(start))
	drawSpike("sooky.png", 500, 500)

	//dbinstance.Write("fookin_data", strings.Join(grammer, " "))

}
