package output

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/input"
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

	g := gol.NewRandomGol(rows, cols, randomSeed)

	golo := NewGolOutputer(g)
	golo.SaveToFile(outputFilePath)

	gr := input.NewGolReader(new(gol.Gol))
	readG, readError := gr.ReadGolFromTextFile(outputFilePath)
	if readError != nil {
		fmt.Errorf("Couldn't load the file %s: %s", outputFilePath, readError)
	}

	if !readG.Equals(g) {
		t.Errorf("Both game of life instances should be equal and they don't")
	}
}
