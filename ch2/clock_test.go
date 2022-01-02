package clock_test

import (
	"bytes"
	"clock"
	"strconv"
	"testing"
	"time"
)

func TestPrintTimeInXminPastY(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	p := clock.Printer{
		Output: fakeTerminal,
	}
	p.PrintTime()
	now := time.Now()
	want := "It's " + strconv.Itoa(now.Minute()) + " minutes past " + strconv.Itoa(now.Hour()) + "\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

}
