package gol

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

// MakeGifAnimation : make a gif animation for some generations
func MakeGifAnimation(g *Gol, outputFilepath string, generations int, delay int) error {
	outputFile, outputFileError := os.Create(outputFilepath)
	if outputFileError != nil {
		return outputFileError
	}

	palette := []color.Color{color.White, color.Black}
	rows := g.grid.rows
	cols := g.grid.cols
	nframes := generations
	gifAnimation := gif.GIF{LoopCount: 0}
	for i := 0; i < nframes; i++ {
		grid := g.grid

		// TODO: Gol method to return an image in output.go
		rect := image.Rect(0, 0, cols, rows)
		frameImage := image.NewPaletted(rect, palette)
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				cell := grid.get(i, j)
				frameImage.SetColorIndex(j, i, uint8(cell))
			}
		}

		gifAnimation.Delay = append(gifAnimation.Delay, delay)
		gifAnimation.Image = append(gifAnimation.Image, frameImage)

		g = g.NextGeneration()
	}
	gif.EncodeAll(outputFile, &gifAnimation)
	return nil
}
