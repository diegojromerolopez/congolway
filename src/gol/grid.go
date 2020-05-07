package gol

import (
	"math/rand"
)

// Status of cells in the Game of Life grid
const (
	ALIVE = 1
	DEAD  = 0
	A     = 1 // ALIVE alias
	D     = 0 // DEAD alias
	VOID  = -1
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
		grid.cells[i] = DEAD
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

func (g *Grid) randomize(randomSeed int64) {
	statuses := []int{ALIVE, DEAD}
	statusesLen := len(statuses)
	for i := 0; i < g.rows*g.cols; i++ {
		g.cells[i] = statuses[rand.Intn(statusesLen)]
	}
}

func (g *Grid) getPos(i int, j int) int {
	return i*g.cols + j
}

func (g *Grid) get(i int, j int) int {
	if i < 0 || i >= g.rows || j < 0 || j >= g.cols {
		return VOID
	}
	pos := g.getPos(i, j)
	return g.cells[pos]
}

func (g *Grid) set(i int, j int, value int) {
	pos := g.getPos(i, j)
	g.cells[pos] = value
}

func (g *Grid) neighbors(i int, j int) []int {
	return []int{
		g.get(i-1, j-1), g.get(i-1, j), g.get(i-1, j+1),
		g.get(i, j-1), g.get(i, j+1),
		g.get(i+1, j-1), g.get(i+1, j), g.get(i+1, j+1),
	}
}

func (g *Grid) aliveNeighborsCount(i int, j int) int {
	aliveCount := 0
	for _, neighborhood := range g.neighbors(i, j) {
		if neighborhood == ALIVE {
			aliveCount++
		}
	}
	return aliveCount
}

func (g *Grid) equals(other *Grid) bool {
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

func (g *Grid) clone() *Grid {
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
