package input

import (
	"fmt"
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/gol"
)

func TestNewGolFromCellsFile5x5(t *testing.T) {
	var expectedCells [][]int = [][]int{
		{D, D, D, D, D},
		{A, A, A, A, A},
		{A, A, D, A, A},
		{A, A, A, A, A},
		{D, D, D, D, D},
	}
	filename := "5x5.cells"
	name := "A 5x5 game of life"
	description := "A 5x5 game of life in .cells format. It is not a known pattern, this is only used for tests."
	testNewGolFromCellsFile(t, filename, name, description, DefaultGeneration, 5, 5,
		DefaultRowLimitation, DefaultColLimitation, DefaultRules, expectedCells)
}

func testNewGolFromCellsFile(t *testing.T, filename string, name string, description string, generation int,
	rows int, cols int, rowLimitation string, colLimitation string, rules string, expectedCells [][]int) {

	g, error := readCellsFile(filename, name, description, generation, rows, cols, rowLimitation, colLimitation, rules)
	if error != nil {
		t.Error(error)
		return
	}
	assertGolIsRight(t, filename, name, description, rows, cols,
		rowLimitation == "limited", colLimitation == "limited", generation, expectedCells, g)
}

func readCellsFile(filename string, name string, description string, generation int,
	rows int, cols int, rowsLimitation string, colsLimitation string, rules string) (base.GolInterface, error) {
	dataFilePath, dataFilePathError := base.GetTestdataFilePath(filename)
	if dataFilePathError != nil {
		return nil, dataFilePathError
	}

	gr := NewGolReader(new(gol.Gol))
	gol, golReadError := gr.ReadCellsFile(dataFilePath)
	if golReadError != nil {
		return nil, fmt.Errorf("Couldn't load the file %s: %s", dataFilePath, golReadError)
	}
	return gol, nil
}
