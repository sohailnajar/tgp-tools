package filefoo_test

import (
	"filefoo"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWriteToFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/write_test_file.txt"
	want := []byte{1, 2, 3}
	err := filefoo.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	perm := stat.Mode().Perm()
	if perm != 0600 {
		t.Errorf("want file mode 0600, got 0%o", perm)
	}

	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}

}

func TestWriteToFileClobbers(t *testing.T) {
	t.Parallel()

	path := t.TempDir() + "/clobber_test.txt"
	err := os.WriteFile(path, []byte{4, 5, 6}, 0600)
	if err != nil {
		t.Fatal(err)
	}
	// overwrite
	want := []byte{1, 2, 3}
	err = filefoo.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}

}

// Test that an existing file does not have open perms
func TestPermsClosed(t *testing.T) {
	t.Parallel()

	path := t.TempDir() + "/perm_test.txt"

	err := os.WriteFile(path, []byte{}, 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = filefoo.WriteToFile(path, []byte{1, 2, 4})
	if err != nil {
		t.Fatal(err)
	}

	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	perm := stat.Mode().Perm()
	if perm != 0600 {
		t.Errorf("want file mode 0600, got 0%o", perm)
	}

}

// test write zeros to named file
func TestWriteZeros(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/zeros_file.txt"
	err := filefoo.BufferdWrite(path, 1000)
	if err != nil {
		t.Fatal(err)
	}
	stat, err2 := os.Stat(path)
	if err2 != nil {
		t.Fatal(err2)
	}
	size := stat.Size()
	if size != 1000 {
		t.Errorf("wanted size 10, got :%d", size)
	}
}

func BenchmarkWriteZeros(b *testing.B) {
	path := b.TempDir() + "/zeros_file.txt"
	for i := 0; i < b.N; i++ {
		filefoo.BufferdWrite(path, 1000)
	}
}
