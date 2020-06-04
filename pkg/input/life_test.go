package input

import (
	"fmt"
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/gol"
)

func TestNewGolFromLife105File(t *testing.T) {
	var expectedCells [][]int = [][]int{
		{A, A, A, D, D, D, D, D, D},
		{A, D, A, D, D, D, D, D, D},
		{D, A, D, D, D, D, D, D, D},
		{D, D, D, D, D, D, D, D, D},
		{D, D, D, D, D, D, A, D, A},
		{D, D, D, D, D, D, A, D, A},
		{D, D, D, D, D, D, A, A, D},
	}
	filename := "3x3_life105.life"
	name := filename
	_, filepathError := base.GetTestdataFilePath(filename)
	if filepathError != nil {
		t.Error(filepathError)
		return
	}

	description := fmt.Sprintf("3x3 life")
	testNewGolFromLifeFile(t, filename, name, description, base.DefaultGeneration, 7, 9,
		base.DefaultRowLimitation, base.DefaultColLimitation, base.DefaultRules, expectedCells)
}

func TestNewGolFromLife106File5x5(t *testing.T) {
	var expectedCells [][]int = [][]int{
		{D, D, D, D, D},
		{A, A, A, A, A},
		{A, A, D, A, A},
		{A, A, A, A, A},
		{A, D, D, D, A},
	}
	filename := "5x5.life"
	name := filename
	filepath, filepathError := base.GetTestdataFilePath(filename)
	if filepathError != nil {
		t.Error(filepathError)
		return
	}

	description := fmt.Sprintf("File path: %s", filepath)
	testNewGolFromLifeFile(t, filename, name, description, base.DefaultGeneration, 5, 5,
		base.DefaultRowLimitation, base.DefaultColLimitation, base.DefaultRules, expectedCells)
}

func testNewGolFromLifeFile(t *testing.T, filename string, name string, description string, generation int,
	rows int, cols int, rowLimitation string, colLimitation string, rules string, expectedCells [][]int) {

	g, error := readLifeFile(filename, name, description, generation, rows, cols, rowLimitation, colLimitation, rules)
	if error != nil {
		t.Error(error)
		return
	}

	assertGolIsRight(t, filename, name, description, rows, cols,
		rowLimitation == "limited", colLimitation == "limited", generation, expectedCells, g)
}

func readLifeFile(filename string, name string, description string, generation int,
	rows int, cols int, rowsLimitation string, colsLimitation string, rules string) (base.GolInterface, error) {
	dataFilePath, dataFilePathError := base.GetTestdataFilePath(filename)
	if dataFilePathError != nil {
		return nil, dataFilePathError
	}

	gr := NewGolReader(new(gol.Gol))
	gol, golReadError := gr.ReadLifeFile(dataFilePath, nil)
	if golReadError != nil {
		return nil, fmt.Errorf("Couldn't load the file %s: %s", dataFilePath, golReadError)
	}
	return gol, nil
}
