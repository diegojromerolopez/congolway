package input

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/animator"
	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/gol"
)

func TestNewGolFromGifFile(t *testing.T) {
	gifFilePath, gifFilePathError := base.GetTestdataFilePath("still.gif")
	if gifFilePathError != nil {
		t.Error(gifFilePathError)
		return
	}

	gi, error := readGifFileFromPath(gifFilePath)
	if error != nil {
		t.Error(error)
		return
	}
	g := gi.(*gol.Gol)

	tempDir, tempDirError := ioutil.TempDir("", "")
	if tempDirError != nil {
		t.Error(tempDirError)
		return
	}
	defer os.RemoveAll(tempDir)

	outputGifPath := filepath.Join(tempDir, "out.gif")
	animator.MakeGif(g, outputGifPath, 1, 0)

	expectedGifBytes, expectedGifErr := ioutil.ReadFile(gifFilePath)
	if expectedGifErr != nil {
		t.Error(expectedGifErr)
		return
	}

	newGifBytes, newGifErr := ioutil.ReadFile(outputGifPath)
	if newGifErr != nil {
		t.Error(newGifErr)
		return
	}

	if !bytes.Equal(expectedGifBytes, newGifBytes) {
		t.Errorf("Not equal")
	}
}

func readGifFileFromPath(filePath string) (base.GolInterface, error) {
	gr := NewGolReader(new(gol.Gol))
	gol, golReadError := gr.ReadGifFile(filePath)
	if golReadError != nil {
		return nil, fmt.Errorf("Couldn't load the file %s: %s", filePath, golReadError)
	}
	return gol, nil
}
