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
	g, error := readGolFromTextFile("10x10.txt")
	if error != nil {
		t.Error(error)
		return
	}

	expectedGifFilePath, _ := base.GetTestdataFilePath("10x10.gif")
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
	MakeGif(g.(*gol.Gol), gifOutputPath, generations, delay)

	gifContents, _ := ioutil.ReadFile(gifOutputPath)
	if bytes.Compare(expectedGifContents, gifContents) != 0 {
		t.Errorf("Should be equal")
	}
}

func readGolFromTextFile(filename string) (base.GolInterface, error) {
	dataFilePath, dataFilePathError := base.GetTestdataFilePath(filename)
	if dataFilePathError != nil {
		return nil, dataFilePathError
	}

	gr := input.NewGolReader(new(gol.Gol))
	gol, golReadError := gr.ReadGolFromTextFile(dataFilePath)
	if golReadError != nil {
		return nil, fmt.Errorf("Couldn't load the file %s: %s", dataFilePath, golReadError)
	}
	return gol, nil
}
