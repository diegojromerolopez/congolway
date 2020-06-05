package input

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
	maxInt := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	minInt := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	file, fileError := os.Open(filepath)
	defer file.Close()

	if fileError != nil {
		return nil, fileError
	}

	file.Seek(0, io.SeekStart)
	reader := bufio.NewReader(file)

	// Get minimum grid size
	maxX := math.MinInt32
	maxY := math.MinInt32
	minX := math.MaxInt32
	minY := math.MaxInt32
	eof := false
	description := ""
	rules := ""
	maxRows := 0
	maxBlockRows := 0
	maxWidth := -1
	for !eof {
		line, lineError := reader.ReadString('\n')
		if lineError != nil {
			if lineError == io.EOF {
				eof = true
			} else {
				return nil, lineError
			}
		}
		if line[0:1] == "#" {
			if line[1:2] == "D" {
				description += strings.TrimSuffix(line[3:], "\n")
			}
			if line[1:2] == "N" {
				rules += base.DefaultRules
			}
			if line[1:2] == "R" {
				rules += line[2:]
			}
			if line[1:2] == "P" {
				maxRows = maxInt(maxRows, maxBlockRows)
				maxBlockRows = 0
				lineParts := strings.Split(strings.TrimSuffix(line, "\n"), " ")
				x, xError := strconv.Atoi(lineParts[1])
				if xError != nil {
					return nil, xError
				}
				y, yError := strconv.Atoi(lineParts[2])
				if yError != nil {
					return nil, yError
				}
				maxX = maxInt(x, maxX)
				minX = minInt(x, minX)
				maxY = maxInt(y, maxY)
				minY = minInt(y, minY)
			}
		} else {
			maxWidth = maxInt(len(strings.TrimSuffix(line, "\n")), maxWidth)
			maxBlockRows++
		}
	}
	maxRows = maxInt(maxRows, maxBlockRows)

	rows := maxY - minY + maxRows
	cols := maxX - minX + maxWidth

	xOffset := -minX
	yOffset := -minY

	if gconf == nil {
		gconf = base.NewDefaultGolConf()
	}
	filepathParts := strings.Split(filepath, "/")
	name := filepathParts[len(filepathParts)-1]
	g := gr.readGol
	g.InitFromConf(name, description, rows, cols, gconf)

	file.Seek(0, io.SeekStart)
	reader = bufio.NewReader(file)

	eof = false
	line, lineError := reader.ReadString('\n')
	for !eof {
		if lineError != nil {
			if lineError == io.EOF {
				eof = true
			} else {
				return nil, lineError
			}
		}
		if line[0:1] == "#" && line[1:2] == "P" {
			trimmedLine := strings.TrimSuffix(line, "\n")
			lineParts := strings.Split(trimmedLine, " ")
			xBlockOffset, xError := strconv.Atoi(lineParts[1])
			if xError != nil {
				return nil, xError
			}
			yBlockOffset, yError := strconv.Atoi(lineParts[2])
			if yError != nil {
				return nil, yError
			}
			rowIndex := yOffset + yBlockOffset
			for !eof {
				blockLine, blockLineError := reader.ReadString('\n')
				if blockLineError != nil {
					if blockLineError == io.EOF {
						eof = true
					} else {
						return nil, blockLineError
					}
				}
				if blockLine[0:1] == "#" {
					line = blockLine
					lineError = blockLineError
					break
				} else {
					colOffset := xOffset + xBlockOffset
					gr.addLife105Row(rowIndex, colOffset, blockLine)
					rowIndex++
				}
			}
		} else {
			line, lineError = reader.ReadString('\n')
		}
	}
	return g, nil
}

// ReadLife106File : read a Game of life from a Life 1.06 file.
// See the following link for more information:
// - 1.06: https://www.conwaylife.com/wiki/Life_1.06
func (gr *GolReader) ReadLife106File(filepath string, gconf *base.GolConf) (base.GolInterface, error) {
	// Set the file reader offset to where the cell positions start
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

func (gr *GolReader) addLife105Row(rowIndex, colOffset int, rawRow string) {
	row := strings.TrimSuffix(rawRow, "\n")
	g := gr.readGol
	for j := 0; j < len(row); j++ {
		if row[j:j+1] == "*" {
			g.Set(rowIndex, colOffset+j, statuses.ALIVE)
		}
	}
}
