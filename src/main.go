package main

import (
	"flag"
	"fmt"

	"github.com/diegojromerolopez/congoay/src/gol"
)

func main() {
	congoayFilePath := flag.String("congoayFilePath", "./samples/oscilator_5x5_gen_0.txt", "Congay file")
	outputFilePath := flag.String("outputFilePath", "out.gif", "File path where the output gif will be saved")

	flag.Parse()

	g, gError := gol.ReadGolFromTextFile(*congoayFilePath)
	if gError != nil {
		fmt.Println(gError.Error())
		return
	}

	gol.MakeGifAnimation(g, *outputFilePath, 50)
}
