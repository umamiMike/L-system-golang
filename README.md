# Implementing L-systems in Go
[L-system - Wikipedia](https://en.wikipedia.org/wiki/L-system)

Having worked with some plant generators in the past, I wanted to understand
some of how they work.  This is a little ongoing experiment in understanding

# You will need
I am working in version 1.11.0.  I have not tested it on earlier versions.

# To run
1. clone the repo  `git clone https://github.com/umamiMike/L-system-golang.git`
2. cd in to the repo `cd L-system-golang`
3. run `go get` 
4. run `make && lsys run -i {integer num of interations}`
5. play with the rules var in rules.go

Currently it only prints the grammer to stdout.

