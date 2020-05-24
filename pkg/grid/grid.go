package grid

import (
	"fmt"
	"math/rand"

	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// Grid : a cell grid implemented as a dense matrix
type Grid struct {
	cells     CellsStorer
	i         func(int) int
	iIsOut    func(int) bool
	j         func(int) int
	jIsOut    func(int) bool
	limitRows bool
	limitCols bool
}

// NewGrid : creates a grid
func NewGrid(rows, cols int, rowLimitation, colLimitation, cellsStorerType string) *Grid {
	cs := CellsStorerFactory(rows, cols, cellsStorerType)
	return newGridFromCellsStorer(rowLimitation, colLimitation, cs)
}

// NewRandomGrid : creates a grid
func NewRandomGrid(rows, cols int, rowLimitation, colLimitation, cellsStorerType string, ramdomSeed int64) *Grid {
	cs := CellsStorerFactory(rows, cols, cellsStorerType)
	grid := newGridFromCellsStorer(rowLimitation, colLimitation, cs)
	grid.Randomize(ramdomSeed)
	return grid
}

// Rows : return the number of rows of the grid
func (g *Grid) Rows() int {
	return g.cells.Rows()
}

// Cols : return the number of columns of the grid
func (g *Grid) Cols() int {
	return g.cells.Cols()
}

// LimitRows : inform if rows are limited or isn't
func (g *Grid) LimitRows() bool {
	return g.limitRows
}

// LimitRowsString : inform if rows are limited or isn't
// by returning the string "limited" or "unlimited"
func (g *Grid) LimitRowsString() string {
	if g.limitRows {
		return "limited"
	}
	return "unlimited"
}

// SetLimitRows : limit or not limit rows.
// If rows are not limited, it will be a circular-by-rows grid
// i.e. if unlimited by rows, on reaching the rows + i column,
// the ith row will be returned
func (g *Grid) SetLimitRows(limitRows bool) {
	rows := g.Rows()
	g.limitRows = limitRows
	if g.limitRows {
		g.i = func(i int) int { return i }
		g.iIsOut = func(i int) bool { return i < 0 || i >= rows }
	} else {
		g.i = func(i int) int { return ((i % rows) + rows) % rows }
		g.iIsOut = func(_ int) bool { return false }
	}
}

// LimitCols : return the number of columns of the grid
func (g *Grid) LimitCols() bool {
	return g.limitCols
}

// LimitColsString : inform if cols are limited or isn't
// by returning the string "limited" or "unlimited"
func (g *Grid) LimitColsString() string {
	if g.limitCols {
		return "limited"
	}
	return "unlimited"
}

// SetLimitCols : limit or not limit cols.
// If cols are not limited, it will be a circular-by-cols grid.
// i.e. if unlimited by columns, on reaching the cols + i column,
// the ith column will be returned
func (g *Grid) SetLimitCols(limitCols bool) {
	cols := g.Cols()
	g.limitCols = limitCols
	if g.limitCols {
		g.j = func(j int) int { return j }
		g.jIsOut = func(j int) bool { return j < 0 || j >= cols }
	} else {
		g.j = func(j int) int { return ((j % cols) + cols) % cols }
		g.jIsOut = func(_ int) bool { return false }
	}
}

// Get : get the value of the cell (ALIVE, DEAD)
//	in the i, j coordinates
func (g *Grid) Get(i, j int) int {
	if g.iIsOut(i) || g.jIsOut(j) {
		return statuses.VOID
	}
	actualI := g.i(i)
	actualJ := g.j(j)
	return g.cells.Get(actualI, actualJ)
}

// Set : set the value of the cell in the i, j coordinates
func (g *Grid) Set(i, j, value int) {
	actualI := g.i(i)
	actualJ := g.j(j)
	g.cells.Set(actualI, actualJ, value)
}

// SetAll : set a value to all ceels
func (g *Grid) SetAll(value int) {
	g.cells.SetAll(value)
}

// Equals : inform if two grids have the same cell value
// for each position.
func (g *Grid) Equals(other *Grid, mode string) bool {
	return g.EqualsError(other, mode) == nil
}

// EqualsError : inform if two grids have the same dimensions and
// the same cell values for each position.
func (g *Grid) EqualsError(other *Grid, mode string) error {
	var cellsEqualsError error
	if mode == "values" {
		cellsEqualsError = g.cells.EqualValuesError(other.cells)
	} else {
		cellsEqualsError = g.cells.EqualsError(other.cells)
	}
	if cellsEqualsError != nil {
		return cellsEqualsError
	}
	if g.limitRows != other.limitRows {
		return fmt.Errorf("Row limits are different: %s vs %s", g.LimitRowsString(), other.LimitRowsString())
	}
	if g.limitCols != other.limitCols {
		return fmt.Errorf("Cols are different: %s vs %s", g.LimitColsString(), other.LimitColsString())
	}
	return nil
}

// Clone : clone the grid in a new grid
func (g *Grid) Clone() *Grid {
	gridClone := newGridFromCellsStorer(g.LimitRowsString(), g.LimitColsString(), g.cells.Clone())
	return gridClone
}

// CloneEmpty : create a new grid with the same size but empty
func (g *Grid) CloneEmpty() *Grid {
	gridEmptyClone := newGridFromCellsStorer(g.LimitRowsString(), g.LimitColsString(), g.cells.CloneEmpty())
	return gridEmptyClone
}

// Randomize : set each cell of the grid to a random (uniform) function
//	according to randomSeed
func (g *Grid) Randomize(randomSeed int64) {
	statusesList := []int{statuses.ALIVE, statuses.DEAD}
	statusesListLen := len(statusesList)
	rows := g.Rows()
	cols := g.Cols()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			g.Set(i, j, statusesList[rand.Intn(statusesListLen)])
		}
	}
}

// NewGridFromCellsStorer : creates a grid
func newGridFromCellsStorer(rowLimitation, colLimitation string, cells CellsStorer) *Grid {
	g := new(Grid)
	if cells == nil {
		panic(fmt.Sprintf("cells argument cannot be nil"))
	}
	g.cells = cells.Clone()
	g.SetLimitRows(rowLimitation == "limited")
	g.SetLimitCols(colLimitation == "limited")
	return g
}

// NewRandomGridFromCellsStorer : creates a randomized grid
func NewRandomGridFromCellsStorer(rowLimitation, colLimitation string, cells CellsStorer, ramdomSeed int64) *Grid {
	grid := newGridFromCellsStorer(rowLimitation, colLimitation, cells)
	grid.Randomize(ramdomSeed)
	return grid
}
