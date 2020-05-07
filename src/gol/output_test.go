package gol

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSaveToTextFile(t *testing.T) {
	testSaveToTextFile(t, 5, 10, int64(1))
	testSaveToTextFile(t, 10, 5, int64(1))
	testSaveToTextFile(t, 10, 10, int64(1))
}

func testSaveToTextFile(t *testing.T, rows int, cols int, randomSeed int64) {
	file, err := ioutil.TempFile("", "temp_gol.txt")
	if err != nil {
		t.Error(err)
		return
	}
	outputFilePath := file.Name()
	defer os.Remove(outputFilePath)

	g := NewRandomGol(rows, cols, randomSeed)
	g.SaveToFile(outputFilePath)

	readG, readError := ReadGolFromTextFile(outputFilePath)
	if readError != nil {
		t.Error(readError)
		return
	}

	if !readG.Equals(g) {
		t.Errorf("Both game of life instances should be equal and they don't")
	}
}
