package testmod

import (
	"os"
	"strings"
)

func Contains(filePath, substring string) (bool, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	return strings.Contains(string(fileContent), substring), nil
}
