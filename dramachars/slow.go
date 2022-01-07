package dramachars

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type printer struct {
	Input  io.Reader
	Output io.Writer
	Delay  time.Duration
}

type option func(printer) printer

func NewPrinter(opts ...option) printer {
	np := printer{
		Input:  os.Stdin,
		Output: os.Stdout,
	}
	for _, opt := range opts {
		np = opt(np)
	}
	return np
}

func WithInput(input io.Reader) option {
	return func(p printer) printer {
		p.Input = input
		return p
	}
}

func WithOutput(output io.Writer) option {
	return func(p printer) printer {
		p.Output = output
		return p
	}
}

func (p printer) PrintSlow() {
	sc := bufio.NewScanner(p.Input)
	sc.Split(bufio.ScanBytes)
	for sc.Scan() {
		time.Sleep(p.Delay)
		fmt.Fprintf(p.Output, "%s", sc.Text())

	}

}
