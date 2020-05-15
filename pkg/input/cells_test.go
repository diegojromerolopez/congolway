package input

import (
	"fmt"
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/neighborhood"
)

func TestNewGolFromCellsFile5x5(t *testing.T) {
	var expectedCells [][]int = [][]int{
		{D, D, D, D, D},
		{A, A, A, A, A},
		{A, A, D, A, A},
		{A, A, A, A, A},
		{D, D, D, D, D},
	}
	testNewGolFromCellsFile(t, "5x5.cells", 5, 5, true, true, 0, expectedCells)
}

func testNewGolFromCellsFile(t *testing.T, filename string, rows int, cols int, limitRows bool, limitCols bool,
	generation int, expectedCells [][]int) {

	limitations := map[bool]string{true: "limited", false: "unlimited"}

	g, error := readCellsFile(filename, rows, cols, limitations[limitRows], limitations[limitCols], generation)
	if error != nil {
		t.Error(error)
		return
	}
	assertGolIsRight(t, filename, rows, cols, limitRows, limitCols, generation, expectedCells, g)
}

func readCellsFile(filename string, rows int, cols int, rowsLimitation string, colsLimitation string, generation int) (base.GolInterface, error) {
	dataFilePath, dataFilePathError := base.GetTestdataFilePath(filename)
	if dataFilePathError != nil {
		return nil, dataFilePathError
	}

	gr := NewGolReader(new(gol.Gol))
	gol, golReadError := gr.ReadCellsFile(dataFilePath, rowsLimitation, colsLimitation, generation, neighborhood.MOORE)
	if golReadError != nil {
		return nil, fmt.Errorf("Couldn't load the file %s: %s", dataFilePath, golReadError)
	}
	return gol, nil
}
