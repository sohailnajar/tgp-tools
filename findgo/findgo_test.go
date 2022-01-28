package findgo_test

import (
	"archive/zip"
	"findgo"
	"os"
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

func TestFilesOnDisk(t *testing.T) {
	t.Parallel()
	fsys := os.DirFS("testdata/findgo")
	want := 4
	got := findgo.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func BenchmarkFilesOnDisk(b *testing.B) {
	fsys := os.DirFS("testdata/findgo")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findgo.Files(fsys)
	}
}

func BenchmarkFilesInMemory(b *testing.B) {
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findgo.Files(fsys)
	}

}

func TestFilesInZip(t *testing.T) {
	t.Parallel()
	fsys, err := zip.OpenReader("testdata/findgo.zip")
	if err != nil {
		t.Fatal(err)
	}
	want := 4
	got := findgo.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
