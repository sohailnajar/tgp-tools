package linecount

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

type counter struct {
	input     io.Reader
	output    io.Writer
	match     string
	wordCount bool
	verbosity bool
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

func WithArgs(args []string) option {
	return func(c *counter) error {
		// see whats passed at cmd
		fset := flag.NewFlagSet(os.Args[0],
			flag.ContinueOnError)
		wordCount := fset.Bool("w", false,
			"Count words instead of lines")
		verbosity := fset.Bool("v", false, "Show verbose count")
		fset.SetOutput(c.output)
		err := fset.Parse(args)
		if err != nil {
			return err
		}
		c.wordCount = *wordCount
		c.verbosity = *verbosity
		// get non-flag arguments
		args = fset.Args()
		if len(args) < 1 {
			return nil
		}
		// args[0] is first non flag argument
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

// Wrapper
func LineCount() int {
	c, err := NewCounter(
		WithArgs(os.Args[1:]),
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

// Wrapper
func Words() int {
	c, err := NewCounter(
		WithArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return c.Words()
}

func RunCli() {
	c, err := NewCounter(
		WithArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if c.wordCount {
		fmt.Println(c.Words())
	} else {
		fmt.Println(c.LineCount())
	}

}
