package animator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/input"
)

func TestMakeGif(t *testing.T) {
	testMakeGif(t, "10x10.txt", "10x10.gif")
	testMakeGif(t, "10x10_unlimited_dims.txt", "10x10_unlimited_dims.gif")
}

func testMakeGif(t *testing.T, filename string, expectedGifFilename string) {
	g, error := readCongolwayFile(filename)
	if error != nil {
		t.Error(error)
		return
	}

	expectedGifFilePath, _ := base.GetTestdataFilePath(expectedGifFilename)
	expectedGifContents, _ := ioutil.ReadFile(expectedGifFilePath)

	gifOutputFile, err := ioutil.TempFile("", "temp_gol.txt")
	if err != nil {
		t.Error(err)
		return
	}
	defer gifOutputFile.Close()
	gifOutputPath := gifOutputFile.Name()

	generations := 10
	delay := 5
	gifError := MakeGif(g.(*gol.Gol), gifOutputPath, generations, delay)
	if gifError != nil {
		t.Error(gifError)
		return
	}

	gifContents, _ := ioutil.ReadFile(gifOutputPath)
	if bytes.Compare(expectedGifContents, gifContents) != 0 {
		t.Errorf("The gif file %s has unexpected content", filename)
	}
}

func readCongolwayFile(filename string) (base.GolInterface, error) {
	dataFilePath, dataFilePathError := base.GetTestdataFilePath(filename)
	if dataFilePathError != nil {
		return nil, dataFilePathError
	}

	gr := input.NewGolReader(new(gol.Gol))
	gol, golReadError := gr.ReadCongolwayFile(dataFilePath)
	if golReadError != nil {
		return nil, fmt.Errorf("Couldn't load the file %s: %s", dataFilePath, golReadError)
	}
	return gol, nil
}
