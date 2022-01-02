package counter

import (
	"fmt"
	"io"
	"os"
)

type Counter struct {
	Num    int
	Output io.Writer
}

func (c *Counter) Next() {
	fmt.Fprintln(c.Output, c.Num)
	c.Num++
}

func NewCounter() Counter {
	return Counter{
		Output: os.Stdout,
		Num:    0,
	}
}

func (c *Counter) Run() {
	for {
		c.Next()
	}
}
