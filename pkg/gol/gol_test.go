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
	g := NewGol(5, 5, 0)
	for i := 0; i < g.Rows(); i++ {
		for j := 0; j < g.Cols(); j++ {
			if g.Get(i, j) != statuses.DEAD {
				t.Errorf("%d, %d cell should be dead", i, j)
			}
		}
	}
}

func TestEquals(t *testing.T) {
	g1, g1ReadError := readGolFromTextFile("10x10.txt")
	if g1ReadError != nil {
		t.Error(g1ReadError)
	}
	g2, g2ReadError := readGolFromTextFile("10x10.txt")
	if g2ReadError != nil {
		t.Error(g2ReadError)
	}
	if !g1.Equals(g2) {
		t.Errorf("Both game of life instances should be equals")
	}
}

func TestNotEquals(t *testing.T) {
	g1, g1ReadError := readGolFromTextFile("10x10.txt")
	if g1ReadError != nil {
		t.Error(g1ReadError)
	}
	g2, g2ReadError := readGolFromTextFile("5x10.txt")
	if g2ReadError != nil {
		t.Error(g2ReadError)
	}
	if g1.Equals(g2) {
		t.Errorf("Both game of life instances should be different")
	}
}

func TestClone(t *testing.T) {
	originalGol, originalGolReadError := readGolFromTextFile("10x10.txt")
	if originalGolReadError != nil {
		t.Error(originalGolReadError)
	}
	cloneGol := originalGol.Clone()
	if !originalGol.Equals(cloneGol) {
		t.Errorf("Clone game of life instance has failed")
	}
}

func TestNextGeneration(t *testing.T) {
	// Test still-life
	testStillNextGeneration(t, "block.txt")
	testStillNextGeneration(t, "bee-hive.txt")
	testStillNextGeneration(t, "loaf.txt")
	testStillNextGeneration(t, "boat.txt")
	testStillNextGeneration(t, "tub.txt")
	// Test oscilators
	testOscilatorNextGeneration(t, "blinker/gen_0.txt", "blinker/gen_1.txt")
	testOscilatorNextGeneration(t, "beacon/gen_0.txt", "beacon/gen_1.txt")
	testOscilatorNextGeneration(t, "toad/gen_0.txt", "toad/gen_1.txt")
	// Test big grids
	testStandardGridNextGeneration(t, "grid1024x1024.txt", "grid1024x1024_gen1.txt", SERIAL)
	testStandardGridNextGeneration(t, "grid1024x1024.txt", "grid1024x1024_gen1.txt", CPUS)
}

func testStandardGridNextGeneration(t *testing.T, gen0FilePath string, gen1FilePath string, goProcesses int) {
	g0, g0ReadError := readGolFromTextFile(gen0FilePath)
	if g0ReadError != nil {
		t.Error(g0ReadError)
	}
	g1, g1ReadError := readGolFromTextFile(gen1FilePath)
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
	g0, g0ReadError := readGolFromTextFile("still/" + stillFilePath)
	if g0ReadError != nil {
		t.Error(g0ReadError)
	}
	g1 := g0.NextGeneration()
	if !g1.GridEquals(g0) {
		t.Errorf("Still-life does not change after a generation")
	}
}

func testOscilatorNextGeneration(t *testing.T, gen0FilePath string, gen1FilePath string) {
	g0, g0ReadError := readGolFromTextFile("oscilators/" + gen0FilePath)
	if g0ReadError != nil {
		t.Error(g0ReadError)
	}

	g1, g1ReadError := readGolFromTextFile("oscilators/" + gen1FilePath)
	if g1ReadError != nil {
		t.Error(g1ReadError)
	}

	if g0.GridEquals(g1) {
		t.Errorf("Odd oscilator game-of-life generation is wrong. They should be different")
	}

	if !g0.NextGeneration().GridEquals(g1) {
		t.Errorf("Odd oscilator game-of-life generation is wrong. They should be equal (odd generation)")
	}

	if !g1.NextGeneration().GridEquals(g0) {
		t.Errorf("Odd oscilator game-of-life generation is wrong. They should be equal (even generation)")
	}
}

func readGolFromTextFile(filename string) (base.GolInterface, error) {
	currentDir, currentDirError := filepath.Abs(".")
	if currentDirError != nil {
		return nil, currentDirError
	}
	dataFilePath := path.Join(currentDir, "..", "..", "testdata", filename)

	gr := input.NewGolReader(new(Gol))
	g, golReadError := gr.ReadGolFromTextFile(dataFilePath)
	if golReadError != nil {
		return nil, fmt.Errorf("Couldn't load the file %s: %s", dataFilePath, golReadError)
	}
	return g, nil
}
