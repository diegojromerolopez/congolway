package output

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/input"
)

func TestSparseSparseToCongolwayFile(t *testing.T) {
	testSaveToCongolwayFile(t, 5, 10, int64(1), "sparse")
	testSaveToCongolwayFile(t, 10, 5, int64(1), "sparse")
	testSaveToCongolwayFile(t, 10, 10, int64(1), "sparse")
}

func TestDenseSaveToCongolwayFile(t *testing.T) {
	testSaveToCongolwayFile(t, 5, 10, int64(1), "dense")
	testSaveToCongolwayFile(t, 10, 5, int64(1), "dense")
	testSaveToCongolwayFile(t, 10, 10, int64(1), "dense")
}

func testSaveToCongolwayFile(t *testing.T, rows int, cols int, randomSeed int64, fileType string) {
	file, err := ioutil.TempFile("", "temp_gol.txt")
	if err != nil {
		t.Error(err)
		return
	}
	outputFilePath := file.Name()
	defer os.Remove(outputFilePath)

	g := gol.NewRandomGol(rows, cols, randomSeed)

	golo := NewGolOutputer(g)
	golo.SaveToCongolwayFile(outputFilePath, fileType)

	gr := input.NewGolReader(new(gol.Gol))
	readG, readError := gr.ReadCongolwayFile(outputFilePath)
	if readError != nil {
		fmt.Errorf("Couldn't load the file %s: %s", outputFilePath, readError)
		return
	}

	if !readG.Equals(g) {
		t.Errorf("Both game of life instances should be equal and they don't")
	}
}
