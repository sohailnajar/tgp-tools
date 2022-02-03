package findold

import (
	"fmt"
	"io/fs"
	"time"
)

func FindOldFiles(fsys fs.FS, age int) (OldFiles []fs.FileInfo) {
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		now := time.Now()
		age := now.AddDate(0, 0, -age)
		fileInfo, _ := d.Info() // skip err check
		if fileInfo.ModTime().Before(age) {
			OldFiles = append(OldFiles, fileInfo)
		}
		return nil
	})
	return OldFiles

}

func PrintOldFiles(fsys fs.FS, age int) {
	OldFiles := FindOldFiles(fsys, age)
	for _, file := range OldFiles {
		fmt.Println(file.Name())
	}
}
