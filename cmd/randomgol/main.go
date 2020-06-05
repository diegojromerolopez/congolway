package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/output"
)

func main() {
	name := flag.String("name", "Random Gol", "Name of the game of life instance that will be created")
	description := flag.String("description", "", "Description of the game of life instance that will be created")
	outputFilePath := flag.String("outputFilePath", "out.txt",
		"File path where the random grid will be saved (.txt, .cells and life extensions are allowed)")
	rows := flag.Int("rows", 100, "Number of rows of the grid")
	cols := flag.Int("columns", 100, "Number of columns of the grid")
	circularRows := flag.String("circularRows", "yes", "Should the rows be circular (yes) or be limited (no)")
	circularCols := flag.String("circularCols", "yes", "Should the columns be circular (yes) or be limited (no)")
	rules := flag.String("rules", "23/3", "Survival and birth rules")
	randomSeed := flag.Int64("randomSeed", 0, "Random seed")
	outputFormat := flag.String("outputFormat", "",
		"Only used for congolway files (.txt files). File format \"dense\" or \"sparse\"")

	flag.Parse()

	rulesMatch, rulerMatchError := regexp.MatchString(`\d+/\d+`, *rules)
	if rulerMatchError != nil || !rulesMatch {
		fmt.Fprintf(os.Stderr, "argument invalid: -rules\n")
		os.Exit(2)
	}

	gridType := "dok"

	rowLimitation := "limited"
	if *circularRows == "yes" {
		rowLimitation = "unlimited"
	}

	colLimitation := "limited"
	if *circularCols == "yes" {
		colLimitation = "unlimited"
	}

	g := gol.NewRandomGol(*name, *description, *rules, gridType, rowLimitation, colLimitation, *rows, *cols, *randomSeed)
	writer := output.NewGolOutputer(g)
	if *outputFormat != "" {
		writer.SaveToCongolwayFile(*outputFilePath, *outputFormat)
	} else {
		writer.SaveToFile(*outputFilePath)
	}
}
