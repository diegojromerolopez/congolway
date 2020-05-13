package animator

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/kettek/apng"
)

// MakeApng : make an animated-png (apng) for a number of generations
// for a game-of-life instance.
func MakeApng(g *gol.Gol, outputFilepath string, generations int) error {
	tempDir, tempDirError := ioutil.TempDir("", "")
	if tempDirError != nil {
		return tempDirError
	}
	defer os.RemoveAll(tempDir)

	numberOfFrames := generations
	imagePaths := make([]string, 0, numberOfFrames)
	for frameIndex := 0; frameIndex < numberOfFrames; frameIndex++ {
		frameOutputFilepath := filepath.Join(tempDir, fmt.Sprintf("png_%d.png", frameIndex))
		pngError := makePng(g, frameOutputFilepath)
		if pngError != nil {
			return pngError
		}
		imagePaths = append(imagePaths, frameOutputFilepath)
		g = g.NextGeneration().(*gol.Gol)
	}

	animation := apng.APNG{
		Frames: make([]apng.Frame, len(imagePaths)),
	}

	outputFile, outputFileError := os.Create(outputFilepath)
	if outputFileError != nil {
		return outputFileError
	}
	defer outputFile.Close()

	for i, imagePath := range imagePaths {
		apngImage, imageError := os.Open(imagePath)
		if imageError != nil {
			return imageError
		}
		defer apngImage.Close()
		pngImage, pngImageError := png.Decode(apngImage)
		if pngImageError != nil {
			return pngImageError
		}
		animation.Frames[i].Image = pngImage
	}
	// Write APNG to our output file
	apng.Encode(outputFile, animation)

	return nil
}

// makePng : make a png for a generation of the game of life
func makePng(g *gol.Gol, outputFilepath string) error {
	outputFile, outputFileError := os.Create(outputFilepath)
	if outputFileError != nil {
		return outputFileError
	}
	// TODO: fix palette to be a map and don't depend on ALIVE = 1, DEAD = 0
	palette := []color.Color{color.White, color.Black}
	rows := g.Rows()
	cols := g.Cols()
	rect := image.Rect(0, 0, cols, rows)
	pngImage := image.NewPaletted(rect, palette)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cell := g.Get(i, j)
			pngImage.SetColorIndex(j, i, uint8(cell))
		}
	}
	png.Encode(outputFile, pngImage)
	return nil
}
