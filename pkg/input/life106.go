package input

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// ReadLife106File : read a Game of life from a Life 1.06 file.
// See the following link for more information:
// - 1.06: https://www.conwaylife.com/wiki/Life_1.06
func (gr *GolReader) ReadLife106File(filepath string, gconf *base.GolConf) (base.GolInterface, error) {
	file, fileError := os.Open(filepath)
	defer file.Close()

	if fileError != nil {
		return nil, fileError
	}

	reader := newLife106Reader(file)

	// Read the #Life 1.06 header
	reader.readLine()

	// Read the dimensions
	rows, cols, dimsError := reader.readDimensions()
	if dimsError != nil {
		return nil, dimsError
	}

	// Prepare the Gol
	if gconf == nil {
		gconf = base.NewDefaultGolConf()
	}
	filepathParts := strings.Split(filepath, "/")
	name := filepathParts[len(filepathParts)-1]
	description := fmt.Sprintf("File path: %s", filepath)

	g := gr.readGol
	g.InitFromConf(name, description, rows, cols, gconf)

	// Read alive cells
	reader.readGrid(rows, cols, g)
	return g, nil
}

type life106Reader struct {
	fr *fileReader
}

func (r *life106Reader) currentLine() *string {
	return r.fr.CurrentLine()
}

func (r *life106Reader) readLine() error {
	return r.fr.ReadLine()
}

func (r *life106Reader) seekStart() {
	r.fr.SeekStart()
}

func (r *life106Reader) readDimensions() (int, int, error) {
	maxRow := -1
	maxCol := -1
	eof := false
	rowNum := 1
	for !eof {
		lineError := r.readLine()
		if lineError != nil {
			if lineError == io.EOF {
				eof = true
			} else {
				return -1, -1, lineError
			}
		}
		line := *r.currentLine()
		if eof && line == "" {
			return -1, -1, lineError
		}
		lineText := strings.TrimSuffix(line, "\n")
		positions := strings.Split(lineText, " ")
		if len(positions) != 2 {
			return -1, -1, fmt.Errorf("Expected two dimensions on row %d, found line \"%s\"", rowNum, lineText)
		}
		row, rowError := strconv.Atoi(positions[0])
		if rowError != nil {
			return -1, -1, rowError
		}
		if maxRow < row {
			maxRow = row
		}
		col, colError := strconv.Atoi(positions[1])
		if colError != nil {
			return -1, -1, colError
		}
		if maxCol < col {
			maxCol = col
		}
		rowNum++
	}
	return maxRow + 1, maxCol + 1, nil
}

func (r *life106Reader) readGrid(rows, cols int, g base.GolInterface) error {
	// To the top of the file again and read the life 1.06 header
	r.seekStart()
	r.readLine()

	eof := false
	rowNum := 1
	for !eof {
		lineError := r.readLine()
		if lineError != nil {
			if lineError == io.EOF {
				eof = true
			} else {
				return lineError
			}
		}
		line := *r.currentLine()
		lineText := strings.TrimSuffix(line, "\n")
		positions := strings.Split(lineText, " ")
		if len(positions) != 2 {
			return fmt.Errorf("Expected two dimensions on row %d, found line \"%s\"", rowNum, lineText)
		}
		row, rowError := strconv.Atoi(positions[0])
		if rowError != nil {
			return rowError
		}
		col, colError := strconv.Atoi(positions[1])
		if colError != nil {
			return colError
		}
		g.Set(row, col, statuses.ALIVE)
		rowNum++
	}
	return nil
}

func newLife106Reader(file *os.File) *life106Reader {
	return &life106Reader{newFileReader(file)}
}
