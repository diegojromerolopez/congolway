package input

import (
	"bufio"
	"fmt"
	"image/color"
	"image/gif"
	"os"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// ReadGifFile : create a new Game of life from a .gif file
func (gr *GolReader) ReadGifFile(filename string, gconf *base.GolConf) (base.GolInterface, error) {
	file, fileError := os.Open(filename)
	defer file.Close()

	if fileError != nil {
		return nil, fileError
	}

	reader := bufio.NewReader(file)
	gifAnim, gifAnimError := gif.DecodeAll(reader)
	if gifAnimError != nil {
		return nil, gifAnimError
	}
	if len(gifAnim.Image) < 1 {
		return nil, fmt.Errorf("Something went wrong loading %s", filename)
	}

	gifStill := gifAnim.Image[0]

	gifBounds := gifStill.Bounds()
	// Max is not included in the bounds but min is
	rows := gifBounds.Max.Y - gifBounds.Min.Y
	cols := gifBounds.Max.X - gifBounds.Min.X

	if gconf == nil {
		gconf = base.NewDefaultGolConf()
	}

	description := fmt.Sprintf("Read from file %s", filename)
	g := gr.readGol
	g.InitFromConf(filename, description, rows, cols, gconf)

	j := 0
	for x := gifBounds.Min.X; x < gifBounds.Max.X; x++ {
		i := 0
		for y := gifBounds.Min.Y; y < gifBounds.Max.Y; y++ {
			gifCellColor := gifStill.At(x, y)
			rgba64Color := color.RGBA64Model.Convert(gifCellColor)
			if !gr.gifPixelIsDeadCell(rgba64Color) {
				g.Set(i, j, statuses.ALIVE)
			}
			i++
		}
		j++
	}
	return g, nil
}

func (gr *GolReader) gifPixelIsDeadCell(clr color.Color) bool {
	r, g, b, a := clr.RGBA()
	// Transparent color
	if a == uint32(0) {
		return true
	}
	// Solid white
	if r == uint32(65535) &&
		g == uint32(65535) &&
		b == uint32(65535) &&
		a == uint32(65535) {
		return true
	}
	return false
}
