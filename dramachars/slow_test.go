package dramachars_test

import (
	"bytes"
	"dramachars"
	"testing"
)

func TestPrintSlow(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	inputBuf := bytes.NewBufferString("sohail")
	want := inputBuf.String()

	c, _ := dramachars.NewPrinter(
		dramachars.WithInput(inputBuf),
		dramachars.WithOutput(fakeTerminal),
	)
	c.PrintSlow()

	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

}

func TestWithEmptyInputArgs(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	inputBuf := bytes.NewBufferString("wake up neo!")
	want := inputBuf.String()

	p, _ := dramachars.NewPrinter(
		dramachars.WithInput(inputBuf),
		dramachars.WithOutput(fakeTerminal),
		dramachars.WithInputArgs([]string{}),
	)
	p.PrintSlow()
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
