package findold_test

import (
	"findold"
	"testing"
	"testing/fstest"
)

func TestOldFiles(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file1": {},
		"file2": {},
		"file3": {},
	}
	want := 3
	got := findold.FindOldFiles(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
