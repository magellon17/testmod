package testmod

import (
	"bytes"
	"os"
)

// Constains checks if the file contains substring
func Contains(filePath, substring string) (bool, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	return bytes.Contains(content, []byte(substring)), nil
}
