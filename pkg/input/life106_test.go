package input

import (
	"fmt"
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/gol"
)

func TestNewGolFromLife106FileDoesNotExist(t *testing.T) {
	filename := "5x5_this_file_does_not_exist.life"
	filepath, filepathError := base.GetTestdataFilePath(filename)
	if filepathError != nil {
		t.Error(filepathError)
		return
	}
	expectedErrorString := fmt.Sprintf("open %s: no such file or directory", filepath)

	gr := NewGolReader(new(gol.Gol))
	_, golReadError := gr.ReadLife106File(filepath, nil)
	if golReadError == nil {
		t.Errorf("Should have returned an error when loading file %s", filepath)
		return
	}
	golReadErrorString := golReadError.Error()

	if golReadErrorString != expectedErrorString {
		t.Errorf("Expected error should be: \"%s\". Returned %s", expectedErrorString, golReadErrorString)
		return
	}
}

func TestNewGolFromLife106EmptyGrid(t *testing.T) {
	filename := "5x5_v1.06_empty_grid.life"
	filepath, filepathError := base.GetTestdataFilePath(filename)
	if filepathError != nil {
		t.Error(filepathError)
		return
	}

	gr := NewGolReader(new(gol.Gol))
	_, golReadError := gr.ReadLife106File(filepath, nil)
	if golReadError == nil {
		t.Errorf("Should have returned an error when loading file %s", filepath)
		return
	}
	golReadErrorString := golReadError.Error()
	if golReadErrorString != "EOF" {
		t.Errorf("Expected error should be: \"EOF\". Returned %s", golReadErrorString)
		return
	}
}

func TestNewGolFromLife106BadRow(t *testing.T) {
	filename := "5x5_v1.06_bad_row.life"
	filepath, filepathError := base.GetTestdataFilePath(filename)
	if filepathError != nil {
		t.Error(filepathError)
		return
	}

	gr := NewGolReader(new(gol.Gol))
	_, golReadError := gr.ReadLife106File(filepath, nil)
	if golReadError == nil {
		t.Errorf("Should have returned an error when loading file %s", filepath)
		return
	}
	golReadErrorString := golReadError.Error()
	if golReadErrorString != "strconv.Atoi: parsing \"a\": invalid syntax" {
		t.Errorf("Expected error should be: \"strconv.Atoi: parsing \"a\": invalid syntax\". Returned %s", golReadErrorString)
		return
	}
}

func TestNewGolFromLife106BadCol(t *testing.T) {
	filename := "5x5_v1.06_bad_col.life"
	filepath, filepathError := base.GetTestdataFilePath(filename)
	if filepathError != nil {
		t.Error(filepathError)
		return
	}

	gr := NewGolReader(new(gol.Gol))
	_, golReadError := gr.ReadLife106File(filepath, nil)
	if golReadError == nil {
		t.Errorf("Should have returned an error when loading file %s", filepath)
		return
	}
	golReadErrorString := golReadError.Error()
	if golReadErrorString != "strconv.Atoi: parsing \"a\": invalid syntax" {
		t.Errorf("Expected error should be: \"strconv.Atoi: parsing \"a\": invalid syntax\". Returned %s", golReadErrorString)
		return
	}
}

func TestNewGolFromLife106NoCol(t *testing.T) {
	filename := "5x5_v1.06_no_col.life"
	filepath, filepathError := base.GetTestdataFilePath(filename)
	if filepathError != nil {
		t.Error(filepathError)
		return
	}
	expectedErrorString := "Expected two dimensions on row 5, found line \"1\""

	gr := NewGolReader(new(gol.Gol))
	_, golReadError := gr.ReadLife106File(filepath, nil)
	if golReadError == nil {
		t.Errorf("Should have returned an error when loading file %s", filepath)
		return
	}
	golReadErrorString := golReadError.Error()
	if golReadErrorString != expectedErrorString {
		t.Errorf("Expected error should be: \"%s\". Returned %s", expectedErrorString, golReadErrorString)
		return
	}
}