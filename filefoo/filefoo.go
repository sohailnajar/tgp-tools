package filefoo

import (
	"os"
)

func WriteToFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0600)
}
