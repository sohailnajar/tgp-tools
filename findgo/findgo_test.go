package findgo_test

import (
	"findgo"
	"testing"
	"testing/fstest"
)

func TestFiles(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file.go":       {},
		"subf/file1.go": {},
		"sub2/file.go":  {},
	}
	want := 3
	got := findgo.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
