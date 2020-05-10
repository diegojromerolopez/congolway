package input

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"testing"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

const A = statuses.ALIVE
const D = statuses.DEAD

func TestNewGolFromTextFile5x10(t *testing.T) {
	var expectedCells [][]int = [][]int{
		{A, A, D, A, A, D, A, A, A, D},
		{A, A, D, A, A, A, D, A, A, A},
		{A, A, D, A, A, A, A, A, D, A},
		{A, A, A, D, A, A, A, A, A, A},
		{A, A, D, A, A, D, A, A, A, A},
	}
	testNewGolFromTextFile(t, "5x10.txt", 5, 10, 543, expectedCells)
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
	testNewGolFromTextFile(t, "10x5.txt", 10, 5, 345, expectedCells)
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
	testNewGolFromTextFile(t, "10x10.txt", 10, 10, 0, expectedCells)
}

func TestNewGolFromTextFile10x10BadVersion(t *testing.T) {
	fmt.Println(os.Args[0])

	currentDir, currentDirError := filepath.Abs(".")
	if currentDirError != nil {
		t.Error(currentDirError)
		return
	}
	dataFilePath := path.Join(currentDir, "..", "..", "testdata", "10x10_bad_version.txt")

	gr := &GolReader{new(gol.Gol)}
	g, error := gr.ReadGolFromTextFile(dataFilePath)

	expectedError := "Unknonwn version found 999999"

	if error != nil {
		if error.Error() != expectedError {
			t.Errorf("Error actual = %v, and Expected = %v.", error, expectedError)
		}
	}

	if g != nil {
		t.Errorf("Bad input file, shouldn't return a pointer to Grid")
	}
}

func testNewGolFromTextFile(t *testing.T, filename string,
	rows int, cols int, generation int, expectedCells [][]int) {

	g, error := readGolFromTextFile(filename)
	if error != nil {
		t.Error(error)
	}

	if g.Generation() != generation {
		t.Errorf("Loaded generation is wrong, got: %d, must be: %d.", g.Generation(), generation)
	}

	if g.Rows() != rows {
		t.Errorf("Loaded number of rows is wrong, got: %d, must be: %d.", g.Rows(), rows)
		return
	}
	if g.Cols() != cols {
		t.Errorf("Loaded number of cols is wrong, got: %d, must be: %d.", g.Cols(), cols)
		return
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			gIJ := g.Get(i, j)
			expectedIJ := expectedCells[i][j]
			if gIJ != expectedIJ {
				t.Errorf("Invalid cell at %d, %d, got: %d, must be: %d.", i, j, gIJ, expectedIJ)
			}
		}
	}
}

func readGolFromTextFile(filename string) (base.GolInterface, error) {
	currentDir, currentDirError := filepath.Abs(".")
	if currentDirError != nil {
		return nil, currentDirError
	}
	dataFilePath := path.Join(currentDir, "..", "..", "testdata", filename)

	gr := &GolReader{new(gol.Gol)}
	gol, golReadError := gr.ReadGolFromTextFile(dataFilePath)
	if golReadError != nil {
		return nil, fmt.Errorf("Couldn't load the file %s: %s", dataFilePath, golReadError)
	}
	return gol, nil
}
