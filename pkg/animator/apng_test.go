package animator

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/gol"
)

func TestMakeApng(t *testing.T) {
	testMakeApng(t, "10x10.txt", "10x10.apng")
	testMakeApng(t, "10x10_unlimited_dims.txt", "10x10_unlimited_dims.apng")
}

func testMakeApng(t *testing.T, filename string, expectedApngFilename string) {
	g, error := readGolFromTextFile(filename)
	if error != nil {
		t.Error(error)
		return
	}

	expectedApngFilePath, _ := base.GetTestdataFilePath(expectedApngFilename)
	expectedApngContents, _ := ioutil.ReadFile(expectedApngFilePath)

	apngOutputFile, err := ioutil.TempFile("", "temp_gol.txt")
	if err != nil {
		t.Error(err)
		return
	}
	defer apngOutputFile.Close()
	apngOutputPath := apngOutputFile.Name()

	generations := 10
	apngError := MakeApng(g.(*gol.Gol), apngOutputPath, generations)
	if apngError != nil {
		t.Error(apngError)
		return
	}

	apngContents, _ := ioutil.ReadFile(apngOutputPath)
	if bytes.Compare(expectedApngContents, apngContents) != 0 {
		t.Errorf("The apng file %s has unexpected content", filename)
	}
}
