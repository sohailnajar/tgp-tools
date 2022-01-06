package dramachars

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type Printer struct {
	Input  io.Reader
	Output io.Writer
	Delay  time.Duration
}

func NewPrinter() Printer {
	return Printer{
		Input:  os.Stdin,
		Output: os.Stdout,
	}
}
func (p Printer) PrintSlow() {
	sc := bufio.NewScanner(p.Input)
	sc.Split(bufio.ScanBytes)
	for sc.Scan() {
		time.Sleep(p.Delay)
		fmt.Fprintf(p.Output, "%s", sc.Text())

	}

}
