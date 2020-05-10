package grid

import (
	"math/rand"

	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// Grid : a grid where the cells develop
type Grid struct {
	cells []int
	rows  int
	cols  int
}

// NewGrid : creates a grid
func NewGrid(rows int, cols int) *Grid {
	grid := new(Grid)
	grid.cells = make([]int, rows*cols)
	for i := 0; i < rows*cols; i++ {
		grid.cells[i] = statuses.DEAD
	}
	grid.rows = rows
	grid.cols = cols
	return grid
}

// NewRandomGrid : creates a grid
func NewRandomGrid(rows int, cols int, ramdomSeed int64) *Grid {
	grid := NewGrid(rows, cols)
	grid.randomize(ramdomSeed)
	return grid
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

// Pos : get the position in the 1-D array of the i, j coordinates
func (g *Grid) Pos(i int, j int) int {
	return i*g.cols + j
}

// Get : get the value of the cell (ALICE, DEAD)
//	in the i, j coordinates
func (g *Grid) Get(i int, j int) int {
	if i < 0 || i >= g.rows || j < 0 || j >= g.cols {
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
	gridClone.rows = g.rows
	gridClone.cols = g.cols
	cellsLength := g.rows * g.cols
	gridClone.cells = make([]int, cellsLength)
	for i := 0; i < cellsLength; i++ {
		gridClone.cells[i] = g.cells[i]
	}
	return gridClone
}
