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

func TestSaveToCellsFileError(t *testing.T) {
	g := gol.NewRandomGol("Random", "", "23/3", "dok", "limited", "limited", 10, 10, 42)

	golo := NewGolOutputer(g)
	savingError := golo.SaveToCellsFile("/non-existant-path.cells")
	if savingError == nil {
		t.Errorf("A saving error should have been returned when saving to a bad path")
		return
	}

}

func testSaveToCellsFile(t *testing.T, rows int, cols int, randomSeed int64) {
	file, err := ioutil.TempFile("", "temp_gol.cells")
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
	readG, readError := gr.ReadCellsFile(outputFilePath, nil)
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
