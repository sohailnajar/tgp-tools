package core_test

import (
	"bytes"
	"strconv"
	"testing"
	"time"

	"github.com/sohailnajar/ch1/core"
)

func PrintTimeInXminPastY(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	core.PrintTimeTo(fakeTerminal)
	now := time.Now()
	want := "It's " + strconv.Itoa(now.Minute()) + " minutes past " + strconv.Itoa(now.Hour())
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

}
