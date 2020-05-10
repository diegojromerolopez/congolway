package base

import (
	"path"
	"path/filepath"
)

// GetTestdataFilePath : return the test data file path for the
// file name passed as argument.
func GetTestdataFilePath(filename string) (string, error) {
	currentDir, currentDirError := filepath.Abs(".")
	if currentDirError != nil {
		return "", currentDirError
	}
	dataFilePath := path.Join(currentDir, "..", "..", "testdata", filename)
	return dataFilePath, nil
}
