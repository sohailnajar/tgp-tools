package dramachars

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type printer struct {
	input  io.Reader
	output io.Writer
	delay  time.Duration
}

type option func(*printer) error

func NewPrinter(opts ...option) (printer, error) {
	np := printer{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(&np)
		if err != nil {
			return printer{}, err
		}

	}
	return np, nil
}

func WithDelay(delay time.Duration) option {
	return func(p *printer) error {
		p.delay = delay
		return nil
	}

}
func WithInput(input io.Reader) option {
	return func(p *printer) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		p.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(p *printer) error {
		if output == nil {
			return errors.New("no output writer")
		}
		p.output = output
		return nil
	}
}

func WithInputArgs(args []string) option {
	return func(p *printer) error {
		if len(args) == 0 {
			return nil
		}
		f, _ := os.Open(args[0])
		p.input = f
		return nil
	}

}

func (p printer) PrintSlow() {
	sc := bufio.NewScanner(p.input)
	sc.Split(bufio.ScanBytes)
	for sc.Scan() {
		time.Sleep(p.delay)
		fmt.Fprintf(p.output, "%s", sc.Text())

	}

}

func PrintSlow() {
	np, err := NewPrinter(
		WithInputArgs(os.Args[1:]),
		WithDelay(time.Second),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	np.PrintSlow()
}
