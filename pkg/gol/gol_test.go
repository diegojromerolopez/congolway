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
