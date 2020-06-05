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
	inputFilePath := flag.String("inputFilePath", "", "File path of the Congolway (.txt), cells (.cells) or life (.life) file")
	generations := flag.Int("generations", 100, "Number of generations of the cellular automaton")
	delay := flag.Int("delay", 500, "Delay between frames, in milliseconds")

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

	gr := input.NewGolReader(new(gol.Gol))
	gi, gError := gr.ReadFile(*inputFilePath, nil)
	if gError != nil {
		fmt.Println(gError.Error())
		return
	}
	g := gi.(*gol.Gol)
	g.SetProcesses(*procs)

	stdoutError := animator.MakeStdout(g, *generations, *delay)
	if stdoutError != nil {
		fmt.Fprintf(os.Stderr, stdoutError.Error())
		os.Exit(1)
	}
}
