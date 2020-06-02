package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// ReadLifeFile : create a new Game of life from a .life file.
// See the following links for more information:
// - 1.05: https://www.conwaylife.com/wiki/Life_1.05
// - 1.06: https://www.conwaylife.com/wiki/Life_1.06
func (gr *GolReader) ReadLifeFile(filename string, gconf *base.GolConf) (base.GolInterface, error) {
	file, fileError := os.Open(filename)
	defer file.Close()

	if fileError != nil {
		return nil, fileError
	}

	reader := bufio.NewReader(file)
	headerLine, headerLineError := reader.ReadString('\n')
	if headerLineError != nil {
		return nil, headerLineError
	}
	headerText := strings.TrimSuffix(headerLine, "\n")
	if headerText == "#Life 1.05" {
		return gr.ReadLife105File(filename, gconf)
	}
	if headerText == "#Life 1.06" {
		return gr.ReadLife106File(filename, gconf)
	}
	return nil, fmt.Errorf("Invalid header \"%s\" for a Life file", headerText)
}

// ReadLife105File : read a Game of life from a Life 1.05 file.
// See the following link for more information:
// - 1.05: https://www.conwaylife.com/wiki/Life_1.05
func (gr *GolReader) ReadLife105File(filepath string, gconf *base.GolConf) (base.GolInterface, error) {
	return nil, fmt.Errorf("Life 1.05 file format is not supported yet")
}

// ReadLife106File : read a Game of life from a Life 1.06 file.
// See the following link for more information:
// - 1.06: https://www.conwaylife.com/wiki/Life_1.06
func (gr *GolReader) ReadLife106File(filepath string, gconf *base.GolConf) (base.GolInterface, error) {

	seekPositionsStart := func(file *os.File) (*bufio.Reader, error) {
		file.Seek(0, io.SeekStart)
		reader := bufio.NewReader(file)
		_, headerLineError := reader.ReadString('\n')
		if headerLineError != nil {
			return nil, headerLineError
		}
		return reader, nil
	}

	file, fileError := os.Open(filepath)
	defer file.Close()

	if fileError != nil {
		return nil, fileError
	}

	reader, readerError := seekPositionsStart(file)
	if readerError != nil {
		return nil, readerError
	}
	maxRow := -1
	maxCol := -1
	for true {
		line, lineError := reader.ReadString('\n')
		if lineError != nil {
			if lineError == io.EOF {
				break
			}
			return nil, lineError
		}
		lineText := strings.TrimSuffix(line, "\n")
		positions := strings.Split(lineText, " ")
		row, rowError := strconv.Atoi(positions[0])
		if rowError != nil {
			return nil, rowError
		}
		if maxRow < row {
			maxRow = row
		}
		col, colError := strconv.Atoi(positions[1])
		if colError != nil {
			return nil, colError
		}
		if maxCol < col {
			maxCol = col
		}
	}

	if gconf == nil {
		gconf = base.NewDefaultGolConf()
	}
	filepathParts := strings.Split(filepath, "/")
	name := filepathParts[len(filepathParts)-1]
	description := fmt.Sprintf("File path: %s", filepath)
	rows := maxRow + 1
	cols := maxCol + 1
	g := gr.readGol
	g.InitFromConf(name, description, rows, cols, gconf)

	// Read alive cells
	reader, readerError = seekPositionsStart(file)
	if readerError != nil {
		return nil, readerError
	}
	eof := false
	for !eof {
		line, lineError := reader.ReadString('\n')
		if lineError != nil {
			if lineError == io.EOF {
				eof = true
			} else {
				return nil, lineError
			}
		}
		lineText := strings.TrimSuffix(line, "\n")
		positions := strings.Split(lineText, " ")
		row, rowError := strconv.Atoi(positions[0])
		if rowError != nil {
			return nil, rowError
		}
		col, colError := strconv.Atoi(positions[1])
		if colError != nil {
			return nil, colError
		}
		g.Set(row, col, statuses.ALIVE)
	}
	return g, nil
}
