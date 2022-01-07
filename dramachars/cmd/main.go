package main

import (
	"dramachars"
	"os"
	"time"
)

func main() {
	np, _ := dramachars.NewPrinter(
		dramachars.WithInputArgs(os.Args[1:]),
		dramachars.WithDelay(time.Second*2),
	)
	np.PrintSlow()

}
