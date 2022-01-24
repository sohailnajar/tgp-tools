package filefoo

import (
	"os"
)

func WriteToFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0600)
	if err != nil {
		return err
	}
	return os.Chmod(path, 0600)
}

func BufferdWrite(path string, size int) error {
	da := []byte{0}
	f, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	for i := 0; i < size; i++ {
		f.Write(da)
	}

	// wr.Write(da)

	return nil
}
