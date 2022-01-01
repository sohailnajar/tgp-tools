package core

import (
	"fmt"
	"io"
	"time"
)

func PrintTimeTo(f io.Writer) {
	now := time.Now()
	fmt.Fprintf(f, "It's %d minutes past %d", now.Minute(), now.Hour())
}
