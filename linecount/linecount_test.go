package linecount_test

import (
	"bytes"
	"linecount"
	"testing"
)

func TestLines(t *testing.T) {
	t.Parallel()
	matchStr := "hello\nis"
	inputBuf := bytes.NewBufferString("hello\nmy\nname\nis")
	c, err := linecount.NewCounter(
		linecount.WithInput(inputBuf),
		linecount.Match(matchStr),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 2
	got := c.LineCount()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestWithInputArgs(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := linecount.NewCounter(
		linecount.WithInput(inputBuf),
		linecount.WithInputArgs([]string{}),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.LineCount()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
