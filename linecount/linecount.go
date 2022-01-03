package linecount

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type counter struct {
	input  io.Reader
	output io.Writer
}

/*
pattern to use when constructor requires
setup. instead of passing many params
we use varidic params and helper functions
that allow us to pass params when necessary
but also allow setup of sane defaults
*/

type option func(*counter) error

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}

func NewCounter(opts ...option) (counter, error) {
	c := counter{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		// option updates counter in place, thats why we pass by ref
		err := opt(&c)
		if err != nil {
			return counter{}, err
		}
	}
	return c, nil
}

func (c counter) LineCount() int {
	lines := 0
	sc := bufio.NewScanner(c.input)
	for sc.Scan() {
		lines++
	}
	return lines

}

func LineCount() int {
	c, err := NewCounter()
	if err != nil {
		panic("internal error") // panic because user can not fix this error
	}
	return c.LineCount()
}
