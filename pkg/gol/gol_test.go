package gol

import (
	"fmt"
	"path"
	"path/filepath"
	"testing"
)

func TestNewGol(t *testing.T) {
	g := NewGol(5, 5)
	grid := g.grid
	for i := 0; i < grid.rows; i++ {
		for j := 0; j < grid.cols; j++ {
			if grid.get(i, j) != DEAD {
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
}

func testStillNextGeneration(t *testing.T, stillFilePath string) {
	g0, g0ReadError := readGolFromTextFile("still/" + stillFilePath)
	if g0ReadError != nil {
		t.Error(g0ReadError)
	}
	g1 := g0.NextGeneration()
	if !g1.GridEquals(g0) {
		g0.Stdout()
		g1.Stdout()
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

func readGolFromTextFile(filename string) (*Gol, error) {
	currentDir, currentDirError := filepath.Abs(".")
	if currentDirError != nil {
		return nil, currentDirError
	}
	dataFilePath := path.Join(currentDir, "test_resources", filename)

	gol, golReadError := ReadGolFromTextFile(dataFilePath)
	if golReadError != nil {
		return nil, fmt.Errorf("Couldn't load the file %s: %s", dataFilePath, golReadError)
	}
	return gol, nil
}