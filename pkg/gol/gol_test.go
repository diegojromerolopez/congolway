package gol

import (
	"fmt"
	"path"
	"path/filepath"
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/input"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

func TestNewGol(t *testing.T) {
	g := NewGol("TestGol", "", "23/3", "dense", "limited", "limited", 0, 5, 5)
	for i := 0; i < g.Rows(); i++ {
		for j := 0; j < g.Cols(); j++ {
			if g.Get(i, j) != statuses.DEAD {
				t.Errorf("%d, %d cell should be dead", i, j)
			}
		}
	}
}

func TestEquals(t *testing.T) {
	g1, g1ReadError := readCongolwayFile("10x10.txt")
	if g1ReadError != nil {
		t.Error(g1ReadError)
	}
	g2, g2ReadError := readCongolwayFile("10x10.txt")
	if g2ReadError != nil {
		t.Error(g2ReadError)
	}
	if !g1.Equals(g2) {
		t.Errorf("Both game of life instances should be equals")
	}
}

func TestNotEquals(t *testing.T) {
	g1, g1ReadError := readCongolwayFile("10x10.txt")
	if g1ReadError != nil {
		t.Error(g1ReadError)
	}
	g2, g2ReadError := readCongolwayFile("5x10.txt")
	if g2ReadError != nil {
		t.Error(g2ReadError)
	}
	if g1.Equals(g2) {
		t.Errorf("Both game of life instances should be different")
	}
}

func TestClone(t *testing.T) {
	originalGol, originalGolReadError := readCongolwayFile("10x10.txt")
	if originalGolReadError != nil {
		t.Error(originalGolReadError)
	}
	cloneGol := originalGol.Clone()
	if !originalGol.Equals(cloneGol) {
		t.Errorf("Clone game of life instance has failed")
	}
}

func TestStillLifeNextGeneration(t *testing.T) {
	// Test still-life
	testStillNextGeneration(t, "block.txt")
	testStillNextGeneration(t, "bee-hive.txt")
	testStillNextGeneration(t, "loaf.txt")
	testStillNextGeneration(t, "boat.txt")
	testStillNextGeneration(t, "tub.txt")
}

func TestOscilatorNextGeneration(t *testing.T) {
	// Test oscilators
	testOscilatorNextGeneration(t, "blinker/gen_0.txt", "blinker/gen_1.txt")
	testOscilatorNextGeneration(t, "beacon/gen_0.txt", "beacon/gen_1.txt")
	testOscilatorNextGeneration(t, "toad/gen_0.txt", "toad/gen_1.txt")
}

func TestBigGridsNextGeneration(t *testing.T) {
	// Test big grids
	testStandardGridNextGeneration(t, "grid1024x1024.txt", "grid1024x1024_gen1.txt", SERIAL)
	testStandardGridNextGeneration(t, "grid1024x1024.txt", "grid1024x1024_gen1.txt", CPUS)
}

func TestPriorChanges(t *testing.T) {
	g, gError := readCongolwayFile("still/boat.txt")
	if gError != nil {
		t.Error(gError)
		return
	}

	changes := [][]int{
		{0, 0, statuses.ALIVE},
		{0, 4, statuses.ALIVE},
		{4, 0, statuses.ALIVE},
		{4, 4, statuses.ALIVE},
		{1, 1, statuses.DEAD},
		{1, 2, statuses.DEAD},
	}
	changedG := g.ChangeCells(changes)

	expectedG, expectedGError := readCongolwayFile("expected_changed_boat.txt")
	if expectedGError != nil {
		t.Error(expectedGError)
		return
	}

	equalsError := changedG.EqualsError(expectedG)
	if equalsError != nil {
		t.Error(equalsError)
	}
}

func testStandardGridNextGeneration(t *testing.T, gen0FilePath string, gen1FilePath string, goProcesses int) {
	g0, g0ReadError := readCongolwayFile(gen0FilePath)
	if g0ReadError != nil {
		t.Error(g0ReadError)
	}
	g1, g1ReadError := readCongolwayFile(gen1FilePath)
	if g1ReadError != nil {
		t.Error(g1ReadError)
	}

	g0.SetProcesses(goProcesses)
	actualG1 := g0.NextGeneration().(*Gol)

	if g0.Equals(actualG1) {
		t.Errorf("Standard-life should change after a generation")
	}
	if !g1.Equals(actualG1) {
		t.Errorf("Standard-life should be equal than the expected gol")
	}
}

func testStillNextGeneration(t *testing.T, stillFilePath string) {
	g0, g0ReadError := readCongolwayFile("still/" + stillFilePath)
	if g0ReadError != nil {
		t.Error(g0ReadError)
	}
	g1 := g0.NextGeneration()
	if !g1.GridEquals(g0, "values") {
		t.Errorf("Still-life does not change after a generation")
	}
}

func testOscilatorNextGeneration(t *testing.T, gen0FilePath string, gen1FilePath string) {
	g0, g0ReadError := readCongolwayFile("oscilators/" + gen0FilePath)
	if g0ReadError != nil {
		t.Error(g0ReadError)
	}

	g1, g1ReadError := readCongolwayFile("oscilators/" + gen1FilePath)
	if g1ReadError != nil {
		t.Error(g1ReadError)
	}

	if g0.GridEquals(g1, "values") {
		t.Errorf("Odd oscilator game-of-life generation is wrong. They should be different")
	}

	if !g0.NextGeneration().GridEquals(g1, "values") {
		t.Errorf("Odd oscilator game-of-life generation is wrong. They should be equal (odd generation)")
	}

	if !g1.NextGeneration().GridEquals(g0, "values") {
		t.Errorf("Odd oscilator game-of-life generation is wrong. They should be equal (even generation)")
	}
}

func readCongolwayFile(filename string) (base.GolInterface, error) {
	currentDir, currentDirError := filepath.Abs(".")
	if currentDirError != nil {
		return nil, currentDirError
	}
	dataFilePath := path.Join(currentDir, "..", "..", "testdata", filename)

	gr := input.NewGolReader(new(Gol))
	g, golReadError := gr.ReadCongolwayFile(dataFilePath)
	if golReadError != nil {
		return nil, fmt.Errorf("Couldn't load the file %s: %s", dataFilePath, golReadError)
	}
	return g, nil
}
