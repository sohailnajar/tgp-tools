package main

import (
	"bytes"
	"dramachars"
	"os"
)

func main() {
	sp := dramachars.Printer{
		Input:  bytes.NewBufferString("Wake up"),
		Output: os.Stdout,
	}
	sp.PrintSlow()
}
