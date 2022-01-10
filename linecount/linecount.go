package linecount

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
)

type counter struct {
	input  io.Reader
	output io.Writer
	match  string
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

func WithInputArgs(args []string) option {
	return func(c *counter) error {
		if len(args) == 0 {
			return nil
		}
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		c.input = f
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

func Match(match string) option {
	return func(c *counter) error {
		c.match = match
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
	mlines := 0
	sc := bufio.NewScanner(c.input)
	for sc.Scan() {
		lines++
		matched, _ := regexp.MatchString(sc.Text(), c.match)
		if matched {
			mlines++
		}
	}

	if c.match != "" {
		return mlines
	} else {
		return lines
	}

}

func LineCount() int {
	c, err := NewCounter(
		WithInputArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return c.LineCount()
}

func (c counter) Words() int {
	words := 0
	scanner := bufio.NewScanner(c.input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}
	return words
}

func Words() int {
	c, err := NewCounter(
		WithInputArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return c.Words()
}
