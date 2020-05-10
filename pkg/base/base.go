package base

import (
	"github.com/diegojromerolopez/congolway/pkg/grid"
)

// GolInterface : minimal Gol interface.
type GolInterface interface {
	Init(rows int, cols int, generation int, neighborhoodType int, gr *grid.Grid)
	Rows() int
	Cols() int
	Clone() GolInterface
	Get(i int, j int) int
	Set(i int, j int, value int)
	Generation() int
	GridEquals(g GolInterface) bool
	Equals(g GolInterface) bool
	NeighborhoodTypeString() string
	NextGeneration() GolInterface
}
