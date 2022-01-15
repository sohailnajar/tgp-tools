package linecount_test

import (
	"bytes"
	"io"
	"linecount"
	"testing"
)

func TestWords(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("hello\nmy\nname\nis jack")
	c, err := linecount.NewCounter(
		linecount.WithInput(inputBuf),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 5
	got := c.Words()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

}
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
		linecount.WithArgs([]string{}),
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

//test flag parsing.
func TestWordCount(t *testing.T) {
	t.Parallel()
	args := []string{"-w", "testdata/three_lines.txt"}
	c, err := linecount.NewCounter(
		linecount.WithArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 6
	got := c.Words()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestBytesCount(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("123")
	c, err := linecount.NewCounter(
		linecount.WithInput(inputBuf),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.BytesCount()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestBogusFlags(t *testing.T) {
	t.Parallel()
	args := []string{"-bogus"}
	_, err := linecount.NewCounter(
		linecount.WithOutput(io.Discard),
		linecount.WithArgs(args),
	)
	if err == nil {
		t.Fatal("want error on bogus flag, got nil")
	}
}
