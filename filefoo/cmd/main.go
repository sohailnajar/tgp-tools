package main

import (
	"filefoo"
	"flag"
)

func main() {
	var size int
	flag.IntVar(&size, "size", 0, "size of the zeros")
	flag.Parse()

	filefoo.BufferdWrite("zeroes.dat", size)

}
