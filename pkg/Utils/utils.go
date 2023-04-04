package Utils

import (
	"fmt"
	"os"
)

func GetDirectoryList(path string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(fmt.Sprintf("./Static/img%s", path))
	if err != nil {
		return nil, err
	}

	return entries, nil
}
