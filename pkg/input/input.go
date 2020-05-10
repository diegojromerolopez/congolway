package input

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// GolReader : tasked with reading a Game of Life from files
type GolReader struct {
	ReadGol base.GolInterface
}

// ReadGolFromTextFile : create a new Game of life from a text file
func (gr *GolReader) ReadGolFromTextFile(filename string) (base.GolInterface, error) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)

	// Read CONGOLWAY header line
	congolwayHeaderLine, congolwayHeaderLineError := gr.readCongolwayFileLine(reader)
	if congolwayHeaderLineError != nil {
		return nil, congolwayHeaderLineError
	}
	if "CONGOLWAY" != congolwayHeaderLine {
		return nil, fmt.Errorf("CONGOLWAY expected, found %s", congolwayHeaderLine)
	}

	// Read version in header line
	versionLine, versionLineError := gr.readCongolwayFileLine(reader)
	if versionLineError != nil {
		return nil, versionLineError
	}
	versionLineMatch, versionLineMatchError := regexp.MatchString(`version:\s*\d+`, versionLine)
	if versionLineMatchError != nil {
		return nil, versionLineMatchError
	}
	if !versionLineMatch {
		return nil, fmt.Errorf("version: D.D where D are positive integers, found %s", versionLine)
	}
	versionDigitRegex := regexp.MustCompile(`\d+`)
	versionString := versionDigitRegex.FindString(versionLine)
	if versionString == "" {
		return nil, fmt.Errorf("version: D.D where D are positive integers, found %s", versionLine)
	}
	version, versionError := strconv.Atoi(versionString)
	if versionError != nil {
		return nil, versionError
	}

	if version == 1 {
		return gr.readTextFileV1(reader)
	}

	return nil, fmt.Errorf("Unknonwn version found %d", version)
}

func (gr *GolReader) readCongolwayFileLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	trimmedLine := strings.TrimSuffix(line, "\n")
	return trimmedLine, nil
}

func (gr *GolReader) readTextFileV1(reader *bufio.Reader) (base.GolInterface, error) {
	// Read generation in header line
	generationLine, generationLineError := gr.readCongolwayFileLine(reader)
	if generationLineError != nil {
		return nil, generationLineError
	}
	generationLineMatch, generationLineMatchError := regexp.MatchString(`generation:\s*\d+`, generationLine)
	if generationLineMatchError != nil {
		return nil, generationLineMatchError
	}
	if !generationLineMatch {
		return nil, fmt.Errorf("generation: D where D is a positive integer, found %s", generationLine)
	}
	generationDigitRegex := regexp.MustCompile(`\d+`)
	generationString := generationDigitRegex.FindString(generationLine)
	if generationString == "" {
		return nil, fmt.Errorf("generation: D where D is a positive integer, found %s", generationString)
	}
	generation, generationError := strconv.Atoi(generationString)
	if generationError != nil {
		return nil, generationError
	}
	if generation < 0 {
		return nil, fmt.Errorf("generation: D where D is a positive integer, found %s", generationString)
	}

	// Read size in header line
	sizeLine, sizeLineError := gr.readCongolwayFileLine(reader)
	if sizeLineError != nil {
		return nil, sizeLineError
	}
	sizeLineMatch, sizeLineMatchError := regexp.MatchString(`size:\s*\d+[^\d]+\d+`, sizeLine)
	if sizeLineMatchError != nil {
		return nil, sizeLineMatchError
	}
	if !sizeLineMatch {
		return nil, fmt.Errorf("size: DxD where D are positive integers, found %s", sizeLine)
	}
	sizeDigitRegex := regexp.MustCompile(`\d+`)
	dimensions := sizeDigitRegex.FindAllString(sizeLine, -1)
	if len(dimensions) != 2 {
		return nil, fmt.Errorf("size: DxD where D are positive integers, found %s", dimensions)
	}
	rows, rowsError := strconv.Atoi(dimensions[0])
	if rowsError != nil {
		return nil, rowsError
	}
	cols, colsError := strconv.Atoi(dimensions[1])
	if colsError != nil {
		return nil, colsError
	}

	// Read grid: header line
	gridLine, gridLineError := gr.readCongolwayFileLine(reader)
	if gridLineError != nil {
		return nil, gridLineError
	}
	if "grid:" != gridLine {
		return nil, fmt.Errorf("grid: expected, found %s", gridLine)
	}

	gr.ReadGol.Init(rows, cols, generation)
	g := gr.ReadGol
	/*
		g := gol.NewGol(rows, cols, generation)

		for rowI := 0; rowI < rows; rowI++ {
			rowString, err := readCongolwayFileLine(reader)
			if err != nil {
				return nil, err
			}
			for colI := 0; colI < cols; colI++ {
				colIStatus := statuses.ALIVE
				if rowString[colI:colI+1] == " " {
					colIStatus = statuses.DEAD
				}
				g.Set(rowI, colI, colIStatus)
			}
		}*/

	for rowI := 0; rowI < rows; rowI++ {
		rowString, err := gr.readCongolwayFileLine(reader)
		if err != nil {
			return nil, err
		}
		for colI := 0; colI < cols; colI++ {
			colIStatus := statuses.ALIVE
			if rowString[colI:colI+1] == " " {
				colIStatus = statuses.DEAD
			}
			g.Set(rowI, colI, colIStatus)
		}
	}
	return g, nil
}
