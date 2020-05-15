package main

import (
	"flag"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/output"
)

func main() {
	name := flag.String("name", "Random Gol", "Name of the game of life instance that will be created")
	description := flag.String("description", "", "Description of the game of life instance that will be created")
	outputFilePath := flag.String("outputFilePath", "out.txt",
		"File path where the random grid will be saved (only .txt and .cells extensions are allowed)")
	rows := flag.Int("rows", 100, "Number of rows of the grid")
	cols := flag.Int("columns", 100, "Number of columns of the grid")
	randomSeed := flag.Int64("randomSeed", 0, "Random seed")
	outputFormat := flag.String("outputFormat", "",
		"Only used for congolway files (.txt files). File format \"dense\" or \"sparse\"")

	flag.Parse()

	g := gol.NewRandomGol(*name, *description, *rows, *cols, *randomSeed)
	writer := output.NewGolOutputer(g)
	if *outputFormat != "" {
		writer.SaveToCongolwayFile(*outputFilePath, *outputFormat)
	} else {
		writer.SaveToFile(*outputFilePath)
	}
}
