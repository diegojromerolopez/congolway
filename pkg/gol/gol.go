package gol

import (
	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/grid"
	"github.com/diegojromerolopez/congolway/pkg/neighborhood"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// Gol : game of life
type Gol struct {
	grid       *grid.Grid
	generation int
}

// NewGol : creates a game of life
func NewGol(rows int, cols int, generation int) *Gol {
	g := new(Gol)
	g.grid = grid.NewGrid(rows, cols)
	g.generation = generation
	return g
}

// NewRandomGol : creates a new random game of life
func NewRandomGol(rows int, cols int, randomSeed int64) *Gol {
	g := new(Gol)
	g.grid = grid.NewRandomGrid(rows, cols, randomSeed)
	g.generation = 0
	return g
}

func (g *Gol) Init(rows int, cols int, generation int) {
	g.grid = grid.NewGrid(rows, cols)
	g.generation = generation
}

// Generation : return the number of generations passed
func (g *Gol) Generation() int {
	return g.generation
}

// Rows : return the number of rows of the grid
func (g *Gol) Rows() int {
	return g.grid.Rows()
}

// Cols : return the number of columns of the grid
func (g *Gol) Cols() int {
	return g.grid.Cols()
}

// Get : get the value of the cell (ALICE, DEAD)
// in the i, j coordinates
func (g *Gol) Get(i int, j int) int {
	return g.grid.Get(i, j)
}

// Set : set the value of the cell in the i, j coordinates
func (g *Gol) Set(i int, j int, value int) {
	g.grid.Set(i, j, value)
}

// NextGeneration : compute the next generation
func (g *Gol) NextGeneration() base.GolInterface {
	const ALIVE = statuses.ALIVE
	const DEAD = statuses.DEAD
	nextG := g.Clone().(*Gol)
	rows := g.Rows()
	cols := g.Cols()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			aliveNeighborsCount := neighborhood.NeighborsWithStatusCount(g, i, j, ALIVE)
			// Text from Wikipedia: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
			// Any live cell with two or three live neighbors survives.
			// Any dead cell with three live neighbors becomes a live cell.
			// All other live cells die in the next generation. Similarly, all other dead cells stay dead.
			switch g.Get(i, j) {
			case ALIVE:
				if aliveNeighborsCount == 2 || aliveNeighborsCount == 3 {
					nextG.Set(i, j, ALIVE)
				} else {
					nextG.Set(i, j, DEAD)
				}
			case DEAD:
				if aliveNeighborsCount == 3 {
					nextG.Set(i, j, ALIVE)
				} else {
					nextG.Set(i, j, DEAD)
				}
			default:
				panic("Invalid cell status")
			}
		}
	}
	nextG.generation++
	return nextG
}

// Equals : inform if two game of life instances have the same data
func (g *Gol) Equals(o base.GolInterface) bool {
	other := o.(*Gol)
	return g.grid.Equals(other.grid) && g.generation == other.generation
}

// GridEquals : inform if two game of life instances have the same data,
//	ignoring the difference in generations value
func (g *Gol) GridEquals(o base.GolInterface) bool {
	other := o.(*Gol)
	return g.grid.Equals(other.grid)
}

// Clone : clone a game of life instance
func (g *Gol) Clone() base.GolInterface {
	clone := new(Gol)
	clone.generation = g.generation
	clone.grid = g.grid.Clone()
	return clone
}
