package counter_test

import (
	"bytes"
	"counter"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	c := counter.Counter{
		Output: fakeTerminal,
	}
	c.Next()
	want := "0\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
