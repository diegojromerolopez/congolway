package output

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/input"
)

func TestSaveToLife105File(t *testing.T) {
	testSaveToLifeFile(t, 5, 10, "1.05", int64(1))
	testSaveToLifeFile(t, 10, 5, "1.05", int64(1))
	testSaveToLifeFile(t, 10, 10, "1.05", int64(1))
}

func TestSaveToLife106File(t *testing.T) {
	testSaveToLifeFile(t, 5, 10, "1.06", int64(1))
	testSaveToLifeFile(t, 10, 5, "1.06", int64(1))
	testSaveToLifeFile(t, 10, 10, "1.06", int64(1))
}

func testSaveToLifeFile(t *testing.T, rows, cols int, version string, randomSeed int64) {
	file, err := ioutil.TempFile("", "temp_gol.txt")
	if err != nil {
		t.Error(err)
		return
	}
	outputFilePath := file.Name()
	defer os.Remove(outputFilePath)
	//outputFilePath := "./out.life"
	outputFilePathParts := strings.Split(outputFilePath, "/")
	name := outputFilePathParts[len(outputFilePathParts)-1]
	description := fmt.Sprintf("File path: %s", outputFilePath)
	g := gol.NewRandomGol(name, description, "23/3", "dok",
		"limited", "limited", rows, cols, randomSeed)

	golo := NewGolOutputer(g)
	golo.SaveToLifeFile(outputFilePath, version)

	gr := input.NewGolReader(new(gol.Gol))
	readG, readError := gr.ReadLifeFile(outputFilePath, nil)
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
