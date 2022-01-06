package dramachars

import (
	"fmt"
	"io"
)

type Printer struct {
	Input  io.Reader
	Output io.Writer
}

func (p Printer) PrintSlow() {
	fmt.Fprintf(p.Output, "%s", p.Input)
}
