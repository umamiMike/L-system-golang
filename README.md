# Implementing L-systems in Go

This project was started about a year into my coding journey. In my work with visual effects I had occasion to use Having worked with some plant generators in the past, I wanted to understand some of how they work.  This is a little ongoing experiment in understanding

I followed the rules layed out in this article -> [L-system - Wikipedia](https://en.wikipedia.org/wiki/L-system)

# You will need
I am working in version 1.11.0.  I have not tested it on earlier versions.

# To run
1. clone the repo  `git clone https://github.com/umamiMike/L-system-golang.git`
2. cd in to the repo `cd L-system-golang`
3. run `go get` 
5. run ` go build -o lsys && ./lsys -iterations=9 -rules="koch_curve" -axiom="F"`
6. play with the rules var in rules.go

Currently it only prints the grammer to stdout and generates an image (WIP)

