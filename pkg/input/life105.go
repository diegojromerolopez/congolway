package input

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
	"github.com/diegojromerolopez/congolway/pkg/utils"
)

// ReadLife105File : read a Game of life from a Life 1.05 file.
// See the following link for more information:
// - 1.05: https://www.conwaylife.com/wiki/Life_1.05
func (gr *GolReader) ReadLife105File(filepath string, gconf *base.GolConf) (base.GolInterface, error) {
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
				maxRows = utils.MaxInt(maxRows, maxBlockRows)
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
				maxX = utils.MaxInt(x, maxX)
				minX = utils.MinInt(x, minX)
				maxY = utils.MaxInt(y, maxY)
				minY = utils.MinInt(y, minY)
			}
		} else {
			maxWidth = utils.MaxInt(len(strings.TrimSuffix(line, "\n")), maxWidth)
			maxBlockRows++
		}
	}
	maxRows = utils.MaxInt(maxRows, maxBlockRows)

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

func (gr *GolReader) addLife105Row(rowIndex, colOffset int, rawRow string) {
	row := strings.TrimSuffix(rawRow, "\n")
	g := gr.readGol
	for j := 0; j < len(row); j++ {
		if row[j:j+1] == "*" {
			g.Set(rowIndex, colOffset+j, statuses.ALIVE)
		}
	}
}
