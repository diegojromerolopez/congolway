package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/diegojromerolopez/congolway/pkg/animator"
	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/input"
)

func main() {
	inputFilePath := flag.String("inputFilePath", "", "File path of the Congolway (.txt) or cells (.cells) file")
	outputFilePath := flag.String("outputFilePath", "out.gif", "File path where the output gif will be saved")
	outputWidth := flag.Int("outputWitdh", -1, "Width of the output gif image. If -1, this image will not be scaled")
	outputHeight := flag.Int("outputHeight", -1, "Height of the output gif image. If -1, this image will not be scaled")
	generations := flag.Int("generations", 100, "Number of generations of the cellular automaton")
	delay := flag.Int("delay", 5, "Delay between frames, in 100ths of a second")
	procsHelp := fmt.Sprintf(
		"Number of GO processes used to compute generations. By default is %d (use as many as hardware CPUs), "+
			"enter a positive integer to set a custom number of proceses", gol.CPUS,
	)
	procs := flag.Int("procs", gol.CPUS, procsHelp)

	flag.Parse()

	if *inputFilePath == "" {
		fmt.Fprintf(os.Stderr, "argument required: -inputFilePath\n")
		os.Exit(2)
	}
	if *procs != gol.CPUS && *procs != gol.SERIAL && *procs < 0 {
		fmt.Fprintf(os.Stderr, "argument invalid: -procs\n")
		os.Exit(2)
	}

	var scaler *animator.ImgScaler
	if *outputWidth > -1 && *outputHeight > -1 {
		scaler = animator.NewImgScaler(*outputWidth, *outputHeight, "NearestNeighbor")
	} else {
		scaler = nil
	}

	gr := input.NewGolReader(new(gol.Gol))
	gi, gError := gr.ReadFile(*inputFilePath)
	if gError != nil {
		fmt.Println(gError.Error())
		return
	}
	g := gi.(*gol.Gol)
	g.SetProcesses(*procs)

	gifError := animator.MakeGif(g, *outputFilePath, *generations, *delay, scaler)
	if gifError != nil {
		fmt.Fprintf(os.Stderr, gifError.Error())
		os.Exit(1)
	}
}
