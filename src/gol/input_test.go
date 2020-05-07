package gol

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"testing"
)

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
	dataFilePath := path.Join(currentDir, "test_resources", "10x10_bad_version.txt")

	g, error := ReadGolFromTextFile(dataFilePath)

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

	if g.generation != generation {
		t.Errorf("Loaded generation is wrong, got: %d, must be: %d.", g.generation, generation)
	}

	grid := g.grid
	if grid.rows != rows {
		t.Errorf("Loaded number of rows is wrong, got: %d, must be: %d.", grid.rows, rows)
	}
	if grid.cols != cols {
		t.Errorf("Loaded number of cols is wrong, got: %d, must be: %d.", grid.cols, cols)
	}

	for i := 0; i < g.grid.rows; i++ {
		for j := 0; j < g.grid.cols; j++ {
			if grid.get(i, j) != expectedCells[i][j] {
				t.Errorf("Invalid cell at %d, %d, got: %d, must be: %d.",
					i, j, grid.get(i, j), expectedCells[i][j])
			}
		}
	}
}
