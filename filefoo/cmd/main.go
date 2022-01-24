package main

import (
	"filefoo"
	"flag"
)

func main() {
	var size int
	var name string
	flag.IntVar(&size, "size", 0, "size of the zeros")
	flag.StringVar(&name, "name", "zeroes.dat", "name of the file")
	flag.Parse()

	filefoo.BufferdWrite(name, size)

}
