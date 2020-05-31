package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/input"
)

func TestSpawnerMainOK(t *testing.T) {
	inputFilePath := "../../testdata/still/block.txt"
	outputFile, err := ioutil.TempFile("", "*output.txt")
	if err != nil {
		t.Error(err)
		return
	}
	defer outputFile.Close()
	outputFilePath := outputFile.Name()

	os.Args = []string{"runner",
		"--inputFilePath", inputFilePath,
		"--outputFilePath", outputFilePath,
		"--generations", "20",
	}
	main()

	g1r := input.NewGolReader(new(gol.Gol))
	gi1, g1ReadError := g1r.ReadFile(inputFilePath)
	if g1ReadError != nil {
		t.Error(g1ReadError)
		return
	}

	g2r := input.NewGolReader(new(gol.Gol))
	gi2, g2ReadError := g2r.ReadFile(outputFilePath)
	if g2ReadError != nil {
		t.Error(g2ReadError)
		return
	}
	gi1.SetGeneration(20)

	equalsError := gi1.EqualsError(gi2)
	if equalsError != nil {
		t.Error(equalsError)
	}
}
