package main

import (
	"counter"
	"os"
)

func main() {
	c := counter.Counter{
		Output: os.Stdout,
		Num:    1,
	}

	c.Run()
}
