package gol

import (
	"fmt"
	"path"
	"path/filepath"
	"testing"
)

func TestEquals(t *testing.T) {
	g1, g1ReadError := readGolFromTextFile("oscilator_5x5_gen_0.txt")
	if g1ReadError != nil {
		t.Error(g1ReadError)
	}
	g2, g2ReadError := readGolFromTextFile("oscilator_5x5_gen_0.txt")
	if g2ReadError != nil {
		t.Error(g2ReadError)
	}
	if !g1.Equals(g2) {
		t.Errorf("Both game of life instances should be equals")
	}
}

func TestNotEquals(t *testing.T) {
	g1, g1ReadError := readGolFromTextFile("oscilator_5x5_gen_0.txt")
	if g1ReadError != nil {
		t.Error(g1ReadError)
	}
	g2, g2ReadError := readGolFromTextFile("oscilator_5x5_gen_1.txt")
	if g2ReadError != nil {
		t.Error(g2ReadError)
	}
	if g1.Equals(g2) {
		t.Errorf("Both game of life instances should be different")
	}
}

func TestClone(t *testing.T) {
	originalGol, originalGolReadError := readGolFromTextFile("oscilator_5x5_gen_0.txt")
	if originalGolReadError != nil {
		t.Error(originalGolReadError)
	}
	cloneGol := originalGol.Clone()
	if !originalGol.Equals(cloneGol) {
		t.Errorf("Clone game of life instance has failed")
	}
}

func TestNextGeneration(t *testing.T) {
	gGen0, gGen0ReadError := readGolFromTextFile("oscilator_5x5_gen_0.txt")
	if gGen0ReadError != nil {
		t.Error(gGen0ReadError)
	}
	cloneGol := gGen0.Clone()
	cloneGol.NextGeneration()

	gGen1, gGen1ReadError := readGolFromTextFile("oscilator_5x5_gen_1.txt")
	if gGen1ReadError != nil {
		t.Error(gGen1ReadError)
	}
	if gGen1.Equals(cloneGol) {
		t.Errorf("Odd oscilator game-of-life generation is wrong")
	}

	cloneGol.NextGeneration()
	if gGen0.Equals(cloneGol) {
		t.Errorf("Even oscilator game-of-life generations must be equal")
	}
}

func TestNextGeneration2(t *testing.T) {
	g0, g0ReadError := readGolFromTextFile("oscilator_5x5_gen_0.txt")
	if g0ReadError != nil {
		t.Error(g0ReadError)
	}
	g1, g1ReadError := readGolFromTextFile("oscilator_5x5_gen_1.txt")
	if g1ReadError != nil {
		t.Error(g1ReadError)
	}

	if g0.Equals(g1) {
		t.Errorf("Even oscilator game-of-life generations must not be equal on different parity generations")
	}

	g0.NextGeneration()

	if !g0.Equals(g1) {
		t.Errorf("Even oscilator game-of-life generations must be equal on same parity generations")
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
