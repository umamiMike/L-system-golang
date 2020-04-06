package main

import (
	"fmt"
	"time"
)

//var dbinstance = db.NewBadger("./tmp/badger")

func main() {
	start := time.Now()
	sys := System{
		iterations: 10,
		rules:      "algea",
		axiom:      "a",
	}

	grammer := sys.generate()
	//fmt.Print(grammer)
	fmt.Println("\n\n", len(grammer))
	fmt.Println(time.Since(start))
	drawSpike("sooky.png", 500, 500)

	//dbinstance.Write("fookin_data", strings.Join(grammer, " "))

}
