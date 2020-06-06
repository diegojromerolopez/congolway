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
	panic(fmt.Sprintf("Invalid grid type: %s. Only \"dense\" or \"dok\" are accepted as gridType values", gridType))
}

// EqualsError : inform if two grids have the same dimensions and
// the same cell values for each position.
func EqualsError(d, o CellsStorer) error {
	if d.Rows() != o.Rows() {
		return fmt.Errorf("Rows are different: %d vs %d", d.Rows(), o.Rows())
	}
	if d.Cols() != o.Cols() {
		return fmt.Errorf("Cols are different: %d vs %d", d.Cols(), o.Cols())
	}
	for i := 0; i < d.Rows(); i++ {
		for j := 0; j < d.Cols(); j++ {
			if d.Get(i, j) != o.Get(i, j) {
				return fmt.Errorf("Cells at (%d,%d) are different: %d vs %d",
					i, j, d.Get(i, j), o.Get(i, j))
			}
		}
	}
	return nil
}
