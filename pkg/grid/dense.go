package grid

import (
	"fmt"
)

// Dense : a cell grid implemented as a dense matrix
type Dense struct {
	cells []int
	rows  int
	cols  int
}

// NewDense : creates a dense grid
func NewDense(rows int, cols int) *Dense {
	dense := new(Dense)
	dense.cells = make([]int, rows*cols)
	dense.rows = rows
	dense.cols = cols
	return dense
}

// Rows : return the number of rows of the grid
func (d *Dense) Rows() int {
	return d.rows
}

// Cols : return the number of columns of the grid
func (d *Dense) Cols() int {
	return d.cols
}

// Get : get the value of the cell (ALICE, DEAD)
//	in the i, j coordinates
func (d *Dense) Get(i int, j int) int {
	d.assertIndexes(i, j)
	pos := d.pos(i, j)
	return d.cells[pos]
}

// Set : set the value of the cell in the i, j coordinates
func (d *Dense) Set(i int, j int, value int) {
	d.assertIndexes(i, j)
	pos := d.pos(i, j)
	d.cells[pos] = value
}

// SetAll : set a value to all ceels
func (d *Dense) SetAll(value int) {
	for i := 0; i < d.rows*d.cols; i++ {
		d.cells[i] = value
	}
}

// Equals : inform if two grids have the same cell value
// for each position.s
func (d *Dense) Equals(other CellsStorer) bool {
	return d.EqualsError(other) == nil
}

// EqualsError : inform if two grids have the same dimensions and
// the same cell values for each position.
func (d *Dense) EqualsError(o CellsStorer) error {
	return EqualsError(d, o)
}

// EqualValues : check value by value if both
// cell storers have the same values
func (d *Dense) EqualValues(o CellsStorer) bool {
	return d.Equals(o)
}

// EqualValuesError : check value by value if both
// cell storers have the same values. Return an error
// if that's not the case
func (d *Dense) EqualValuesError(o CellsStorer) error {
	return d.EqualsError(o)
}

// Clone : clone the grid in a new grid
func (d *Dense) Clone() CellsStorer {
	gridClone := NewDense(d.rows, d.cols)
	for i := 0; i < d.rows*d.cols; i++ {
		gridClone.cells[i] = d.cells[i]
	}
	return gridClone
}

// CloneEmpty : create a new grid with the same size but empty
func (d *Dense) CloneEmpty() CellsStorer {
	return NewDense(d.rows, d.cols)
}

// assertIndexes : assert the position is legal in the cell storage
func (d *Dense) assertIndexes(i, j int) {
	if i < 0 || i >= d.rows {
		panic(fmt.Sprintf("Invalid row index: %d not in [0, %d]", i, d.rows-1))
	}
	if j < 0 || j >= d.cols {
		panic(fmt.Sprintf("Invalid col index: %d not in [0, %d]", j, d.cols-1))
	}
}

// pos : get the position in the 1-D array of the i, j coordinates
func (d *Dense) pos(i int, j int) int {
	return i*d.cols + j
}
