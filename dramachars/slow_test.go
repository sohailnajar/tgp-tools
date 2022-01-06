package dramachars_test

import (
	"bytes"
	"dramachars"
	"testing"
)

func TestPrintSlow(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	inputBuf := bytes.NewBufferString("Wake up! Neo")
	c := dramachars.Printer{
		Input:  inputBuf,
		Output: fakeTerminal,
	}
	c.PrintSlow()
	want := inputBuf.String()
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

}
