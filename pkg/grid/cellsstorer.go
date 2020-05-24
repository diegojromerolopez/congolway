package grid

import (
	"fmt"
	"strings"

	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// CellsStorer : minimal storage of cells in a grid.
type CellsStorer interface {
	Rows() int
	Cols() int
	Get(i int, j int) int
	Set(i int, j int, value int)
	SetAll(value int)
	Equals(other CellsStorer) bool
	EqualsError(other CellsStorer) error
	EqualValues(other CellsStorer) bool
	EqualValuesError(other CellsStorer) error
	Clone() CellsStorer
	CloneEmpty() CellsStorer
}

// CellsStorerFactory : creates a new grid from type string
func CellsStorerFactory(rows, cols int, gridType string) CellsStorer {
	if strings.ToLower(gridType) == "dense" {
		return NewDense(rows, cols)
	}
	if strings.ToLower(gridType) == "dok" {
		return NewDok(rows, cols, statuses.DEAD)
	}
	panic(fmt.Sprintf("Invalid grid type: %s. Only \"dense\" or \"dok\" ", gridType))
}
