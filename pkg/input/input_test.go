package input

import (
	"fmt"
	"path"
	"path/filepath"

	"testing"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

const A = statuses.ALIVE
const D = statuses.DEAD

func TestNewGolFromTextFile3x3SparseDefaultDead(t *testing.T) {
	var expectedCells [][]int = [][]int{
		{A, D, D},
		{D, A, D},
		{D, D, A},
	}
	testNewGolFromTextFile(t, "3x3_sparse_default_dead.txt", 3, 3, true, true, 543, expectedCells)
}

func TestNewGolFromTextFile3x3SparseDefaultAlive(t *testing.T) {
	var expectedCells [][]int = [][]int{
		{D, A, A},
		{A, D, A},
		{A, A, D},
	}
	testNewGolFromTextFile(t, "3x3_sparse_default_alive.txt", 3, 3, true, true, 543, expectedCells)
}

func TestNewGolFromTextFile5x10(t *testing.T) {
	var expectedCells [][]int = [][]int{
		{A, A, D, A, A, D, A, A, A, D},
		{A, A, D, A, A, A, D, A, A, A},
		{A, A, D, A, A, A, A, A, D, A},
		{A, A, A, D, A, A, A, A, A, A},
		{A, A, D, A, A, D, A, A, A, A},
	}
	testNewGolFromTextFile(t, "5x10.txt", 5, 10, true, true, 543, expectedCells)
}

func TestNewGolFromTextFile10x5(t *testing.T) {
	var expectedCells [][]int = [][]int{
		{A, A, D, A, A},
		{A, A, D, A, A},
		{A, A, D, A, A},
		{A, A, A, D, A},
		{A, A, D, A, A},
		{A, A, A, D, A},
		{A, D, A, D, A},
		{A, A, D, A, A},
		{A, A, D, A, A},
		{D, A, D, A, A},
	}
	testNewGolFromTextFile(t, "10x5.txt", 10, 5, true, true, 345, expectedCells)
}

func TestNewGolFromTextFile10x10(t *testing.T) {
	var expectedCells [][]int = [][]int{
		{A, A, D, A, A, D, A, A, A, D},
		{A, A, D, A, A, A, D, A, A, A},
		{A, A, D, A, A, A, A, A, D, A},
		{A, A, A, D, A, A, A, A, A, A},
		{A, A, D, A, A, D, A, A, A, A},
		{A, A, A, D, A, A, A, D, A, A},
		{A, D, A, D, A, A, D, A, A, A},
		{A, A, D, A, A, A, D, A, D, A},
		{A, A, D, A, A, A, D, A, A, A},
		{D, A, D, A, A, A, D, A, A, D},
	}
	testNewGolFromTextFile(t, "10x10.txt", 10, 10, true, true, 0, expectedCells)
}

func TestNewGolFromTextFile10x10WithLimitedRows(t *testing.T) {
	var expectedCells [][]int = [][]int{
		{A, A, D, A, A, D, A, A, A, D},
		{A, A, D, A, A, A, D, A, A, A},
		{A, A, D, A, A, A, A, A, D, A},
		{A, A, A, D, A, A, A, A, A, A},
		{A, A, D, A, A, D, A, A, A, A},
		{A, A, A, D, A, A, A, D, A, A},
		{A, D, A, D, A, A, D, A, A, A},
		{A, A, D, A, A, A, D, A, D, A},
		{A, A, D, A, A, A, D, A, A, A},
		{D, A, D, A, A, A, D, A, A, D},
	}
	testNewGolFromTextFile(t, "10x10_limited_rows.txt", 10, 10, true, false, 0, expectedCells)
}

func TestNewGolFromTextFile10x10WithLimitedCols(t *testing.T) {
	var expectedCells [][]int = [][]int{
		{A, A, D, A, A, D, A, A, A, D},
		{A, A, D, A, A, A, D, A, A, A},
		{A, A, D, A, A, A, A, A, D, A},
		{A, A, A, D, A, A, A, A, A, A},
		{A, A, D, A, A, D, A, A, A, A},
		{A, A, A, D, A, A, A, D, A, A},
		{A, D, A, D, A, A, D, A, A, A},
		{A, A, D, A, A, A, D, A, D, A},
		{A, A, D, A, A, A, D, A, A, A},
		{D, A, D, A, A, A, D, A, A, D},
	}
	testNewGolFromTextFile(t, "10x10_limited_cols.txt", 10, 10, false, true, 0, expectedCells)
}

func TestNewGolFromTextFile10x10WithUnlimitedDimensions(t *testing.T) {
	var expectedCells [][]int = [][]int{
		{A, A, D, A, A, D, A, A, A, D},
		{A, A, D, A, A, A, D, A, A, A},
		{A, A, D, A, A, A, A, A, D, A},
		{A, A, A, D, A, A, A, A, A, A},
		{A, A, D, A, A, D, A, A, A, A},
		{A, A, A, D, A, A, A, D, A, A},
		{A, D, A, D, A, A, D, A, A, A},
		{A, A, D, A, A, A, D, A, D, A},
		{A, A, D, A, A, A, D, A, A, A},
		{D, A, D, A, A, A, D, A, A, D},
	}
	testNewGolFromTextFile(t, "10x10_unlimited_dims.txt", 10, 10, false, false, 0, expectedCells)
}

func TestNewGolFromTextFile10x10BadVersion(t *testing.T) {
	currentDir, currentDirError := filepath.Abs(".")
	if currentDirError != nil {
		t.Error(currentDirError)
		return
	}
	dataFilePath := path.Join(currentDir, "..", "..", "testdata", "10x10_bad_version.txt")

	gr := NewGolReader(new(gol.Gol))
	g, error := gr.ReadGolFromTextFile(dataFilePath)

	expectedError := "Unknonwn version found 999999"

	if error != nil {
		if error.Error() != expectedError {
			t.Errorf("Error actual = %v, and Expected = %v.", error, expectedError)
			return
		}
	}

	if g != nil {
		t.Errorf("Bad input file, shouldn't return a pointer to Grid")
		return
	}
}

func testNewGolFromTextFile(t *testing.T, filename string,
	rows int, cols int, limitRows bool, limitCols bool,
	generation int, expectedCells [][]int) {

	g, error := readGolFromTextFile(filename)
	if error != nil {
		t.Error(error)
		return
	}

	if g.Generation() != generation {
		t.Errorf("Loaded generation is wrong, got: %d, must be: %d.", g.Generation(), generation)
		return
	}

	if g.NeighborhoodTypeString() != "Moore" {
		t.Errorf("Loaded neighborhood is wrong, got: %s, must be: Moore.", g.NeighborhoodTypeString())
		return
	}

	if g.Rows() != rows {
		t.Errorf("Loaded number of rows is wrong, got: %d, must be: %d.", g.Rows(), rows)
		return
	}
	if g.Cols() != cols {
		t.Errorf("Loaded number of cols is wrong, got: %d, must be: %d.", g.Cols(), cols)
		return
	}

	if g.LimitRows() != limitRows {
		if limitRows {
			t.Errorf("Should limit rows, but it isn't.")
		} else {
			t.Errorf("Shouldn't limit rows, but it is.")
		}
		return
	}
	if g.LimitCols() != limitCols {
		if limitCols {
			t.Errorf("Should limit cols, but it isn't.")
		} else {
			t.Errorf("Shouldn't limit cols, but it is.")
		}
		return
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			gIJ := g.Get(i, j)
			expectedIJ := expectedCells[i][j]
			if gIJ != expectedIJ {
				t.Errorf("Invalid cell at %d,%d. It got: %d, must be: %d.", i, j, gIJ, expectedIJ)
			}
		}
	}
}

func readGolFromTextFile(filename string) (base.GolInterface, error) {
	dataFilePath, dataFilePathError := base.GetTestdataFilePath(filename)
	if dataFilePathError != nil {
		return nil, dataFilePathError
	}

	gr := NewGolReader(new(gol.Gol))
	gol, golReadError := gr.ReadGolFromTextFile(dataFilePath)
	if golReadError != nil {
		return nil, fmt.Errorf("Couldn't load the file %s: %s", dataFilePath, golReadError)
	}
	return gol, nil
}
