package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/diegojromerolopez/congolway/src/gol"
)

func main() {
	inputFilePath := flag.String("inputFilePath", "", "Input Congolway file")
	outputFilePath := flag.String("outputFilePath", "out.gif", "File path where the output gif will be saved")
	generations := flag.Int("generations", 100, "Number of generations of the cellular automaton")
	delay := flag.Int("delay", 5, "Delay between frames, in 100ths of a second")

	flag.Parse()

	if *inputFilePath == "" {
		fmt.Fprintf(os.Stderr, "argument required: -inputFilePath\n")
		os.Exit(2)
	}

	g, gError := gol.ReadGolFromTextFile(*inputFilePath)
	if gError != nil {
		fmt.Println(gError.Error())
		return
	}

	gol.MakeGifAnimation(g, *outputFilePath, *generations, *delay)
}
