package output

import (
	"fmt"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// GolOutputer : tasked with writting in several devices the passed
// Game of Life (gol.Gol object)
type GolOutputer struct {
	gol *gol.Gol
}

// NewGolOutputer : returns a new pointer to GolOutputer
func NewGolOutputer(g *gol.Gol) *GolOutputer {
	return &GolOutputer{g}
}

func (gout *GolOutputer) get(i int, j int) int {
	return gout.gol.Get(i, j)
}

func (gout *GolOutputer) generation() int {
	return gout.gol.Generation()
}

func (gout *GolOutputer) neighborhoodTypeString() string {
	return gout.gol.NeighborhoodTypeString()
}

func (gout *GolOutputer) limitsString() string {
	limitsStr := ""
	if gout.gol.LimitRows() && gout.gol.LimitCols() {
		limitsStr += "rows, cols"
	} else if gout.gol.LimitRows() {
		limitsStr += "rows"
	} else if gout.gol.LimitCols() {
		limitsStr += "cols"
	} else {
		panic("Impossible condition")
	}
	return limitsStr
}

// Stdout : prints on stdout the current state of the grid
func (gout *GolOutputer) Stdout() {
	rows := gout.gol.Rows()
	cols := gout.gol.Cols()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if gout.get(i, j) == statuses.ALIVE {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
