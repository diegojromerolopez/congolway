package gol

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ReadGolFromTextFile : create a new Game of life from a text file
func ReadGolFromTextFile(filename string) (*Gol, error) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)

	// Read CONGOAY header line
	congayHeaderLine, congayHeaderLineError := readCongayFileLine(reader)
	if congayHeaderLineError != nil {
		return nil, congayHeaderLineError
	}
	if "CONGOAY" != congayHeaderLine {
		return nil, fmt.Errorf("CONGOAY expected, found %s", congayHeaderLine)
	}

	// Read version in header line
	versionLine, versionLineError := readCongayFileLine(reader)
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
		return readTextFileV1(reader)
	}

	return nil, fmt.Errorf("Unknonwn version found %d", version)
}

func readCongayFileLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	trimmedLine := strings.TrimSuffix(line, "\n")
	return trimmedLine, nil
}

func readTextFileV1(reader *bufio.Reader) (*Gol, error) {
	// Read generation in header line
	generationLine, generationLineError := readCongayFileLine(reader)
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
	sizeLine, sizeLineError := readCongayFileLine(reader)
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
	gridLine, gridLineError := readCongayFileLine(reader)
	if gridLineError != nil {
		return nil, gridLineError
	}
	if "grid:" != gridLine {
		return nil, fmt.Errorf("grid: expected, found %s", gridLine)
	}

	grid := NewGrid(rows, cols)
	for rowI := 0; rowI < rows; rowI++ {
		rowString, err := readCongayFileLine(reader)
		if err != nil {
			return nil, err
		}
		for colI := 0; colI < cols; colI++ {
			colIStatus := ALIVE
			if rowString[colI:colI+1] == " " {
				colIStatus = DEAD
			}
			grid.set(rowI, colI, colIStatus)
		}
	}
	g := new(Gol)
	g.grid = grid
	g.generation = generation

	return g, nil
}
