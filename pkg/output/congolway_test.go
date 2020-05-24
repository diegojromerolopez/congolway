package output

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/input"
)

func TestSparseSavedToCongolwayFile(t *testing.T) {
	testRandomGolSavedToCongolwayFile(t, 20, 30, "23/3", int64(1), "sparse")
	testRandomGolSavedToCongolwayFile(t, 10, 5, "23/3", int64(1), "sparse")
	testRandomGolSavedToCongolwayFile(t, 10, 10, "23/3", int64(1), "sparse")
}

func TestDenseSavedToCongolwayFile(t *testing.T) {
	testRandomGolSavedToCongolwayFile(t, 5, 10, "23/3", int64(1), "dense")
	testRandomGolSavedToCongolwayFile(t, 10, 5, "23/3", int64(1), "dense")
	testRandomGolSavedToCongolwayFile(t, 10, 10, "23/3", int64(1), "dense")
}

func testRandomGolSavedToCongolwayFile(t *testing.T, rows int, cols int, rules string, randomSeed int64, fileType string) {
	file, err := ioutil.TempFile("", "temp_gol.txt")
	if err != nil {
		t.Error(err)
		return
	}
	outputFilePath := file.Name()
	defer os.Remove(outputFilePath)

	g := gol.NewRandomGol("Random", "", "23/3", "dok", "limited", "limited", rows, cols, randomSeed)

	golo := NewGolOutputer(g)
	golo.SaveToCongolwayFile(outputFilePath, fileType)

	gr := input.NewGolReader(new(gol.Gol))
	readG, readError := gr.ReadCongolwayFile(outputFilePath)
	if readError != nil {
		t.Error(fmt.Errorf("Couldn't load the file %s: %s", outputFilePath, readError))
		return
	}

	equalsError := readG.EqualsError(g)
	if equalsError != nil {
		t.Error(equalsError)
		return
	}

	if !readG.Equals(g) {
		t.Errorf("Both game of life instances should be equal and they don't")
	}
}
