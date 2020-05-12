package animator

import (
	"fmt"
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// MakeSvg : make a gif animation for some generations
func MakeSvg(g *gol.Gol, outputFilepath string, generations int, delay int) error {
	rows := g.Rows()
	cols := g.Cols()

	outputFile, outputFileError := os.Create(outputFilepath)
	if outputFileError != nil {
		return outputFileError
	}
	canvas := svg.New(outputFile)
	canvas.Start(cols, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cellID := fmt.Sprintf("c_%d_%d", i, j)
			if g.Get(i, j) == statuses.ALIVE {
				canvas.Square(j, i, 1, `fill="black"`, fmt.Sprintf(`id="%s"`, cellID))
			} else {
				canvas.Square(j, i, 1, `fill="black"`, `opacity="0"`, fmt.Sprintf(`id="%s"`, cellID))
			}
		}
	}
	earlierG := g
	numberOfFrames := generations
	for frameIndex := 0; frameIndex < numberOfFrames; frameIndex++ {
		animationDelay := delay * frameIndex
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				cellID := fmt.Sprintf("c_%d_%d", i, j)
				cellSelector := fmt.Sprintf("#%s", cellID)
				cellValue := g.Get(i, j)
				if earlierG.Get(i, j) != cellValue {
					if cellValue == statuses.ALIVE {
						canvas.Animate(cellSelector, "opacity", 0, 1, float64(delay), 0, fmt.Sprintf(`begin="%ds"`, animationDelay))
					} else if cellValue == statuses.DEAD {
						canvas.Animate(cellSelector, "opacity", 1, 0, float64(delay), 0, fmt.Sprintf(`begin="%ds"`, animationDelay))
					}
				}
			}
		}
		earlierG = g
		g = g.NextGeneration().(*gol.Gol)
	}
	canvas.End()
	return nil
}
