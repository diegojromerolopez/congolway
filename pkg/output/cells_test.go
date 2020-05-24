package output

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/input"
)

func TestSaveToCellsFile(t *testing.T) {
	testSaveToCellsFile(t, 5, 10, int64(1))
	testSaveToCellsFile(t, 10, 5, int64(1))
	testSaveToCellsFile(t, 10, 10, int64(1))
}

func testSaveToCellsFile(t *testing.T, rows int, cols int, randomSeed int64) {
	file, err := ioutil.TempFile("", "temp_gol.txt")
	if err != nil {
		t.Error(err)
		return
	}
	outputFilePath := file.Name()
	defer os.Remove(outputFilePath)

	g := gol.NewRandomGol("Random", "", "23/3", "dok", "limited", "limited", rows, cols, randomSeed)

	golo := NewGolOutputer(g)
	golo.SaveToCellsFile(outputFilePath)

	gr := input.NewGolReader(new(gol.Gol))
	readG, readError := gr.ReadCellsFile(outputFilePath)
	if readError != nil {
		t.Error(fmt.Errorf("Couldn't load the file %s: %s", outputFilePath, readError))
		return
	}

	equalsError := readG.EqualsError(g)
	if equalsError != nil {
		t.Error(equalsError)
		return
	}

	if !readG.Equals(g) {
		t.Errorf("Both game of life instances should be equal and they don't")
	}
}
