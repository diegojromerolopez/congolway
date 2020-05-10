package animator

import (
	"image"
	"image/color"
	"image/gif"
	"os"

	"github.com/diegojromerolopez/congolway/pkg/gol"
)

// MakeGif : make a gif animation for some generations
func MakeGif(g *gol.Gol, outputFilepath string, generations int, delay int) error {
	outputFile, outputFileError := os.Create(outputFilepath)
	if outputFileError != nil {
		return outputFileError
	}

	palette := []color.Color{color.White, color.Black}
	rows := g.Rows()
	cols := g.Cols()
	nframes := generations
	gifAnimation := gif.GIF{LoopCount: 0}
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, cols, rows)
		frameImage := image.NewPaletted(rect, palette)
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				cell := g.Get(i, j)
				frameImage.SetColorIndex(j, i, uint8(cell))
			}
		}

		gifAnimation.Delay = append(gifAnimation.Delay, delay)
		gifAnimation.Image = append(gifAnimation.Image, frameImage)

		g = g.NextGeneration().(*gol.Gol)
	}
	gif.EncodeAll(outputFile, &gifAnimation)
	return nil
}
