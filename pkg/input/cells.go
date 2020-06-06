package input

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// ReadCellsFile : create a new Game of life from a .cells file
func (gr *GolReader) ReadCellsFile(filename string, gconf *base.GolConf) (base.GolInterface, error) {
	file, fileError := os.Open(filename)
	defer file.Close()

	if fileError != nil {
		return nil, fileError
	}

	reader := newCellsReader(file)

	// Name of the GOL pattern
	name, nameError := reader.readName()
	if nameError != nil {
		return nil, nameError
	}

	// Description of the pattern
	description, descriptionError := reader.readDescription()
	if descriptionError != nil {
		return nil, descriptionError
	}

	// Dimensions of the pattern
	rows, cols, dimsError := reader.readDimensions()
	if dimsError != nil {
		return nil, dimsError
	}

	g := gr.readGol

	if gconf == nil {
		gconf = base.NewDefaultGolConf()
	}

	g.InitFromConf(name, description, rows, cols, gconf)

	gridError := reader.readGrid(rows, cols, g)
	if gridError != nil {
		return nil, gridError
	}

	return g, nil
}

type cellsReader struct {
	fr *fileReader
}

func (r *cellsReader) currentLine() *string {
	return r.fr.CurrentLine()
}

func (r *cellsReader) readLine() error {
	return r.fr.ReadLine()
}

func (r *cellsReader) seekStart() {
	r.fr.SeekStart()
}

func (r *cellsReader) readName() (string, error) {
	nameError := r.readLine()
	if nameError != nil {
		return "", nameError
	}
	name := strings.TrimPrefix(*r.currentLine(), "!Name: ")
	return name, nil
}

func (r *cellsReader) readDescription() (string, error) {
	description := ""
	// Description of the pattern
	for true {
		err := r.readLine()
		if err != nil {
			return "", err
		}
		if (*r.currentLine())[0:1] != "!" {
			break
		} else {
			currentLine := *r.currentLine()
			description += strings.TrimSuffix(currentLine[1:], " ") + " "
		}
	}
	description = strings.TrimSuffix(description, " ")
	return description, nil
}

func (r *cellsReader) readDimensions() (int, int, error) {
	rows := 0
	cols := len(*r.currentLine())
	lastLoop := false
	for true {
		rows++
		if lastLoop {
			break
		}
		err := r.readLine()
		if err != nil {
			if err == io.EOF {
				lastLoop = true
			} else {
				return -1, -1, err
			}
		}
	}
	return rows, cols, nil
}

func (r *cellsReader) readGrid(rows, cols int, g base.GolInterface) error {
	r.seekStart()

	// Name of the GOL pattern
	_, nameError := r.readName()
	if nameError != nil {
		return nameError
	}

	// Description of the pattern
	_, descriptionError := r.readDescription()
	if descriptionError != nil {
		return descriptionError
	}

	// Read grid
	cellValueCorrespondence := map[string]int{".": statuses.DEAD, "O": statuses.ALIVE}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cellIJ := (*r.currentLine())[j : j+1]
			cellValue, cellValueOK := cellValueCorrespondence[cellIJ]
			if !cellValueOK {
				return fmt.Errorf("Value %s in the cell %d,%d is not a valid one. Only \".\" or \"O\" values are allowed", cellIJ, i, j)
			}
			g.Set(i, j, cellValue)
		}
		r.readLine()
	}
	return nil
}

func newCellsReader(file *os.File) *cellsReader {
	return &cellsReader{newFileReader(file)}
}
