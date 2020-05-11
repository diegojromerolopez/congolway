package grid

import (
	"math/rand"

	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// Grid : a grid where the cells develop
type Grid struct {
	cells     []int
	rows      int
	cols      int
	i         func(int) int
	iIsOut    func(int) bool
	j         func(int) int
	jIsOut    func(int) bool
	limitRows bool
	limitCols bool
}

// NewGrid : creates a grid
func NewGrid(rows int, cols int, rowLimitation string, colLimitation string) *Grid {
	g := new(Grid)
	g.init(rows, cols, rowLimitation == "limited", colLimitation == "limited", nil)
	return g
}

// NewRandomGrid : creates a grid
func NewRandomGrid(rows int, cols int, rowLimitation string, colLimitation string, ramdomSeed int64) *Grid {
	grid := NewGrid(rows, cols, rowLimitation, colLimitation)
	grid.randomize(ramdomSeed)
	return grid
}

func (g *Grid) init(rows int, cols int, limitRows bool, limitCols bool, cells []int) {
	g.cells = make([]int, rows*cols)
	if cells == nil {
		for i := 0; i < rows*cols; i++ {
			g.cells[i] = statuses.DEAD
		}
	} else {
		for i := 0; i < rows*cols; i++ {
			g.cells[i] = cells[i]
		}
	}

	g.rows = rows
	g.cols = cols

	g.limitRows = limitRows
	g.limitCols = limitCols

	if g.limitRows {
		g.i = func(i int) int { return i }
		g.iIsOut = func(i int) bool { return i < 0 || i >= rows }
	} else {
		g.i = func(i int) int { return ((i % rows) + rows) % rows }
		g.iIsOut = func(_ int) bool { return false }
	}
	if g.limitCols {
		g.j = func(j int) int { return j }
		g.jIsOut = func(j int) bool { return j < 0 || j >= cols }
	} else {
		g.j = func(j int) int { return ((j % cols) + cols) % cols }
		g.jIsOut = func(_ int) bool { return false }
	}
}

// Randomize : set each cell of the grid to a random (uniform) function
//	according to randomSeed
func (g *Grid) randomize(randomSeed int64) {
	statusesList := []int{statuses.ALIVE, statuses.DEAD}
	statusesListLen := len(statusesList)
	for i := 0; i < g.rows*g.cols; i++ {
		g.cells[i] = statusesList[rand.Intn(statusesListLen)]
	}
}

// Rows : return the number of rows of the grid
func (g *Grid) Rows() int {
	return g.rows
}

// Cols : return the number of columns of the grid
func (g *Grid) Cols() int {
	return g.cols
}

// LimitRows : inform if rows are limited or isn't
func (g *Grid) LimitRows() bool {
	return g.limitRows
}

// LimitCols : return the number of columns of the grid
func (g *Grid) LimitCols() bool {
	return g.limitCols
}

// Pos : get the position in the 1-D array of the i, j coordinates
func (g *Grid) Pos(i int, j int) int {
	actualI := g.i(i)
	actualJ := g.j(j)
	return actualI*g.cols + actualJ
}

// Get : get the value of the cell (ALICE, DEAD)
//	in the i, j coordinates
func (g *Grid) Get(i int, j int) int {
	if g.iIsOut(i) || g.jIsOut(j) {
		return statuses.VOID
	}
	pos := g.Pos(i, j)
	return g.cells[pos]
}

// Set : set the value of the cell in the i, j coordinates
func (g *Grid) Set(i int, j int, value int) {
	pos := g.Pos(i, j)
	g.cells[pos] = value
}

// Equals : inform if two grids have the same cell value
// for each position.s
func (g *Grid) Equals(other *Grid) bool {
	if g.rows != other.rows || g.cols != g.cols {
		return false
	}
	for pos := 0; pos < g.rows*g.cols; pos++ {
		if g.cells[pos] != other.cells[pos] {
			return false
		}
	}
	return true
}

// Clone : clone the grid in a new grid
func (g *Grid) Clone() *Grid {
	gridClone := new(Grid)
	gridClone.init(g.rows, g.cols, g.limitRows, g.limitRows, g.cells)
	return gridClone
}

// CloneEmpty : create a new grid with the same size but empty
func (g *Grid) CloneEmpty() *Grid {
	gridEmptyClone := new(Grid)
	gridEmptyClone.init(g.rows, g.cols, g.limitRows, g.limitRows, nil)
	return gridEmptyClone
}
