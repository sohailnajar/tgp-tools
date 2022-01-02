package clock

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Printer struct {
	Output io.Writer
}

func (p Printer) PrintTime() {
	now := time.Now()
	fmt.Fprintf(p.Output, "It's %d minutes past %d\n", now.Minute(), now.Hour())
}

// Set defaults
func NewPrinter() Printer {
	return Printer{
		Output: os.Stdout,
	}
}

func Print() {
	NewPrinter().PrintTime()
}
