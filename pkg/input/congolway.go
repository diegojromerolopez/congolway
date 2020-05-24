package input

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/neighborhood"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// ReadCongolwayFile : create a new Game of life from a text file
func (gr *GolReader) ReadCongolwayFile(filename string) (base.GolInterface, error) {
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
		return gr.readCongolwayFileV1(reader)
	}

	return nil, fmt.Errorf("Unknonwn version found %d", version)
}

func (gr *GolReader) readCongolwayFileV1(reader *bufio.Reader) (base.GolInterface, error) {
	// Read name in header line
	nameLine, nameLineError := gr.readCongolwayFileLine(reader)
	if nameLineError != nil {
		return nil, nameLineError
	}
	nameLineMatch, nameLineMatchError := regexp.MatchString(`name:\s*.+`, nameLine)
	if nameLineMatchError != nil {
		return nil, nameLineMatchError
	}
	if !nameLineMatch {
		return nil, fmt.Errorf("name: <name of the game of life instance>, found %s", nameLine)
	}
	namePrefixRegex := regexp.MustCompile(`name:\s*(.+)`)
	name := strings.Trim(namePrefixRegex.ReplaceAllString(nameLine, "${1}"), " ")

	// Read description in header line
	descriptionLine, descriptionLineError := gr.readCongolwayFileLine(reader)
	if descriptionLineError != nil {
		return nil, descriptionLineError
	}
	descriptionLineMatch, descriptionLineMatchError := regexp.MatchString(`description:\s*.+`, descriptionLine)
	if descriptionLineMatchError != nil {
		return nil, descriptionLineMatchError
	}
	if !descriptionLineMatch {
		return nil, fmt.Errorf("description: <description of the game of life instance>, found %s", descriptionLine)
	}
	descriptionPrefixRegex := regexp.MustCompile(`description:\s*(.+)`)
	description := strings.Trim(descriptionPrefixRegex.ReplaceAllString(descriptionLine, "${1}"), " ")

	// Rules
	rulesLine, rulesLineError := gr.readCongolwayFileLine(reader)
	if rulesLineError != nil {
		return nil, rulesLineError
	}
	rulesLineMatch, rulesLineMatchError := regexp.MatchString(`rules:\s*\d+/\d+`, rulesLine)
	if rulesLineMatchError != nil {
		return nil, rulesLineMatchError
	}
	if !rulesLineMatch {
		return nil, fmt.Errorf("rules: survival/birth, found %s", rulesLine)
	}
	rulesPrefixRegex := regexp.MustCompile(`\d+/\d+`)
	rulesMatches := rulesPrefixRegex.FindAllString(rulesLine, -1)
	if len(rulesMatches) != 1 {
		return nil, fmt.Errorf("rules: survival/birth invalid match. Found %s", rulesLine)
	}
	rules := rulesMatches[0]

	// Generation
	generationOccurences, generationError := gr.readTextFileLine(
		reader, regexp.MustCompile(`generation:\s*\d+`), regexp.MustCompile(`\d+`), 1,
	)
	if generationError != nil {
		return nil, generationError
	}
	generation, generationError := strconv.Atoi(generationOccurences[0])
	if generationError != nil {
		return nil, generationError
	}
	if generation < 0 {
		return nil, fmt.Errorf("generation: D where D is a positive integer, found %d", generation)
	}

	// Read neighborhood type:
	neighborhoodType := neighborhood.NONE
	neighLine, neighLineError := gr.readCongolwayFileLine(reader)
	if neighLineError != nil {
		return nil, neighLineError
	}
	if neighLine == "neighborhood_type: Moore" {
		neighborhoodType = neighborhood.MOORE
	} else if neighLine == "neighborhood type: Von Neumann" {
		neighborhoodType = neighborhood.VONNEUMANN
	} else {
		return nil, fmt.Errorf("\"neighborhood_type: Moore\" or \"neighborhood_type: Von Neumman\", found %s", neighLine)
	}

	// Read dimensions of the grid
	dimensions, dimensionsError := gr.readTextFileLine(
		reader, regexp.MustCompile(`size:\s*\d+[^\d]+\d+`), regexp.MustCompile(`\d+`), 2,
	)
	if dimensionsError != nil {
		return nil, dimensionsError
	}
	rows, rowsError := strconv.Atoi(dimensions[0])
	if rowsError != nil {
		return nil, rowsError
	}
	cols, colsError := strconv.Atoi(dimensions[1])
	if colsError != nil {
		return nil, colsError
	}

	limits, limitsError := gr.readTextFileLine(
		reader, regexp.MustCompile(`limits:\s*(rows)?,?\s*(cols)?`), nil, -1,
	)
	if limitsError != nil {
		return nil, limitsError
	}
	rowsLimitationRegex := regexp.MustCompile(`rows`)
	rowLimitationMatches := rowsLimitationRegex.FindAllString(limits[0], -1)
	rowsLimitation := "no"
	if len(rowLimitationMatches) > 0 {
		rowsLimitation = "limited"
	}
	colsLimitationRegex := regexp.MustCompile(`cols`)
	colsLimitationMatches := colsLimitationRegex.FindAllString(limits[0], -1)
	colsLimitation := "no"
	if len(colsLimitationMatches) > 0 {
		colsLimitation = "limited"
	}

	gr.readGol.Init(name, description, generation,
		rows, cols, rowsLimitation, colsLimitation,
		rules, neighborhoodType)

	// Read grid type
	gridTypeLine, gridTypeLineError := gr.readCongolwayFileLine(reader)
	if gridTypeLineError != nil {
		return nil, gridTypeLineError
	}
	gritTypeLineParts := strings.Split(gridTypeLine, " ")
	if len(gritTypeLineParts) != 2 {
		fmt.Errorf("\"grid_type: dense\" or \"grid_type: exparse\" expected, found %s", gridTypeLine)
	}
	gridType := gritTypeLineParts[1]
	if gridType != "dense" && gridType != "sparse" {
		return nil, fmt.Errorf("Invalid grid_type. Only dense and sparse values are accepted, found %s", gridType)
	}

	// Read grid: header line
	gridLine, gridLineError := gr.readCongolwayFileLine(reader)
	if gridLineError != nil {
		return nil, gridLineError
	}
	if "grid:" != gridLine {
		return nil, fmt.Errorf("grid: expected, found %s", gridLine)
	}

	if gridType == "dense" {
		// TODO: read X as 1 and space as 0
		return gr.readGridInDenseFormat(reader)
	}
	if gridType == "sparse" {
		// TODO: read number of status
		statusesCount := 2
		return gr.readGridInSparseFormat(reader, statusesCount)
	}
	return nil, fmt.Errorf("Invalid grid_type. Only dense and sparse values are accepted, found %s", gridType)
}

func (gr *GolReader) readTextFileLine(
	reader *bufio.Reader, lineMatcher *regexp.Regexp,
	findRegex *regexp.Regexp, mandatoryFoundCount int,
) ([]string, error) {
	// Line read from file
	line, lineError := gr.readCongolwayFileLine(reader)
	if lineError != nil {
		return nil, lineError
	}

	// No match required, return the line
	if lineMatcher == nil {
		return []string{line}, nil
	}

	// Check if matches desired regex
	lineMatch := lineMatcher.MatchString(line)
	if !lineMatch {
		return nil, fmt.Errorf("%s does not match the desired regex %s", line, lineMatcher.String())
	}

	// No find regex present, return the line
	if findRegex == nil {
		return []string{line}, nil
	}

	// Find all strings occurences, making sure the number of occurrences
	// (len(foundOccurrences)) is what it must to be (mandatoryFoundCount)
	foundOccurrences := findRegex.FindAllString(line, -1)
	if mandatoryFoundCount < 0 {
		return foundOccurrences, nil
	}
	if len(foundOccurrences) != mandatoryFoundCount {
		return nil, fmt.Errorf("%s does not contain %d match for the regex %s", line, mandatoryFoundCount, findRegex.String())
	}
	return foundOccurrences, nil
}

func (gr *GolReader) readCongolwayFileLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	trimmedLine := strings.TrimSuffix(line, "\n")
	return trimmedLine, nil
}

func (gr *GolReader) readGridInDenseFormat(reader *bufio.Reader) (base.GolInterface, error) {
	g := gr.readGol
	rows := g.Rows()
	cols := g.Cols()
	for rowI := 0; rowI < rows; rowI++ {
		rowString, err := gr.readCongolwayFileLine(reader)
		if err != nil {
			return nil, err
		}
		for colI := 0; colI < cols; colI++ {
			colIStatus := statuses.ALIVE
			cellValue := rowString[colI : colI+1]
			if cellValue == " " || cellValue == "0" {
				colIStatus = statuses.DEAD
			}
			g.Set(rowI, colI, colIStatus)
		}
	}
	return g, nil
}

func (gr *GolReader) readGridInSparseFormat(reader *bufio.Reader, numberOfStatus int) (base.GolInterface, error) {
	g := gr.readGol

	defaultLine, defaultLineError := gr.readCongolwayFileLine(reader)
	if defaultLineError != nil {
		return nil, defaultLineError
	}
	defaultParts := strings.Split(defaultLine, ":")
	if len(defaultParts) != 2 {

	}
	defaultTitle := defaultParts[0]
	if defaultTitle != "default" {
		return nil, fmt.Errorf("Expected default status value, found %s", defaultTitle)
	}
	defaultStatusValue, defaultStatusValueError := strconv.Atoi(strings.TrimSpace(defaultParts[1]))
	if defaultStatusValueError != nil {
		return nil, fmt.Errorf("Invalid default value, found %d", defaultStatusValue)
	}
	g.SetAll(defaultStatusValue)

	for statusI := 0; statusI < numberOfStatus; statusI++ {
		rowStringI, rowStringIError := gr.readCongolwayFileLine(reader)
		if rowStringIError != nil {
			return nil, rowStringIError
		}
		// TODO: check status is valid
		status, coords, lineError := sparseLineToCoordinates(rowStringI)
		if lineError != nil {
			return nil, lineError
		}
		for _, coord := range coords {
			g.Set(coord.i, coord.j, status)
		}
	}
	return g, nil
}

type gridCell struct {
	i int
	j int
}

func sparseLineToCoordinates(sparseLine string) (int, []gridCell, error) {
	sparseLineParts := strings.Split(sparseLine, ":")
	// TODO: check this error
	status, _ := strconv.Atoi(sparseLineParts[0])
	coordinatesString := sparseLineParts[1]
	if coordinatesString == "" {
		return status, nil, nil
	}

	pointsRegexp := regexp.MustCompile(`\(\d+,\s*\d+\)`)
	coordinates := make([]gridCell, 0, 100)
	for _, coordinateString := range pointsRegexp.FindAllString(coordinatesString, -1) {
		ij := strings.Split(coordinateString[1:len(coordinateString)-1], ",")
		// TODO: check these errors
		i, _ := strconv.Atoi(ij[0])
		j, _ := strconv.Atoi(ij[1])
		coordinates = append(coordinates, gridCell{i, j})
	}
	return int(status), coordinates, nil
}
