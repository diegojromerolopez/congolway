package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/input"
	"github.com/diegojromerolopez/congolway/pkg/output"
)

func main() {
	inputFilePath := flag.String("inputFilePath", "", "File path of the Congolway (.txt/.congol) or cells (.cells) file")
	outputFilePath := flag.String("outputFilePath", "", "File path of the output Congolway (.txt/.congol) or cells (.cells) file")

	flag.Parse()

	if *inputFilePath == "" {
		fmt.Fprintf(os.Stderr, "argument required: -inputFilePath\n")
		os.Exit(2)
	}

	if *outputFilePath == "" {
		fmt.Fprintf(os.Stderr, "argument required: -outputFilePath\n")
		os.Exit(2)
	}

	if *inputFilePath == *outputFilePath {
		fmt.Fprintf(os.Stderr, "inputFilePath and outputFilePath cannot be the equal\n")
		os.Exit(2)
	}

	gr := input.NewGolReader(new(gol.Gol))
	gi, gError := gr.ReadFile(*inputFilePath)
	if gError != nil {
		fmt.Println(gError.Error())
		return
	}
	g := gi.(*gol.Gol)
	writer := output.NewGolOutputer(g)
	writer.SaveToFile(*outputFilePath)
}
