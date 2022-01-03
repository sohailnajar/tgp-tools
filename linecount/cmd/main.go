package main

import (
	"fmt"
	"linecount"
)

func main() {
	lc := linecount.LineCount()
	fmt.Println(lc)
}
