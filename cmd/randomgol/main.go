package main

import (
	"flag"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/output"
)

func main() {
	outputFilePath := flag.String("outputFilePath", "out.txt", "File path where the random grid will be saved")
	rows := flag.Int("rows", 100, "Number of rows of the grid")
	cols := flag.Int("columns", 100, "Number of columns of the grid")
	randomSeed := flag.Int64("randomSeed", 0, "Random seed")
	outputFormat := flag.String("outputFormat", "dense", "File format \"dense\" or \"sparse\"")

	flag.Parse()

	g := gol.NewRandomGol(*rows, *cols, *randomSeed)
	writer := output.NewGolOutputer(g)
	writer.SaveToFile(*outputFilePath, *outputFormat)
}
