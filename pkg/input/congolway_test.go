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
	g, error := gr.ReadCongolwayFile(dataFilePath)

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

func testNewGolFromTextFile(t *testing.T, filename string, rows int, cols int, limitRows bool, limitCols bool,
	generation int, expectedCells [][]int) {

	g, error := readCongolwayFile(filename)
	if error != nil {
		t.Error(error)
		return
	}
	assertGolIsRight(t, filename, rows, cols, limitRows, limitCols, generation, expectedCells, g)
}

func readCongolwayFile(filename string) (base.GolInterface, error) {
	dataFilePath, dataFilePathError := base.GetTestdataFilePath(filename)
	if dataFilePathError != nil {
		return nil, dataFilePathError
	}

	gr := NewGolReader(new(gol.Gol))
	gol, golReadError := gr.ReadCongolwayFile(dataFilePath)
	if golReadError != nil {
		return nil, fmt.Errorf("Couldn't load the file %s: %s", dataFilePath, golReadError)
	}
	return gol, nil
}
