package gol

import (
	"fmt"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/grid"
	"github.com/diegojromerolopez/congolway/pkg/neighborhood"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// Gol : game of life
type Gol struct {
	grid             *grid.Grid
	generation       int
	neighborhoodType int
	neighborhoodFunc neighborhood.Func
}

// NewGol : creates a game of life
func NewGol(rows int, cols int, generation int) *Gol {
	g := new(Gol)
	gr := grid.NewGrid(rows, cols, "limited", "limited")
	g.InitWithGrid(generation, neighborhood.MOORE, gr)
	return g
}

// NewRandomGol : creates a new random game of life
func NewRandomGol(rows int, cols int, randomSeed int64) *Gol {
	g := new(Gol)
	gr := grid.NewRandomGrid(rows, cols, "limited", "limited", randomSeed)
	g.InitWithGrid(0, neighborhood.MOORE, gr)
	return g
}

// Init : initialize a Game of Life instance
func (g *Gol) Init(rows int, cols int, rowsLimitation string, colsLimitation string, generation int, neighborhoodType int) {
	g.grid = grid.NewGrid(rows, cols, rowsLimitation, colsLimitation)
	g.generation = generation
	g.neighborhoodType = neighborhoodType
	g.neighborhoodFunc = neighborhood.GetFunc(g.neighborhoodType)
}

// InitWithGrid : initialize a Game of Life instance
func (g *Gol) InitWithGrid(generation int, neighborhoodType int, gr *grid.Grid) {
	g.grid = gr
	g.generation = generation
	g.neighborhoodType = neighborhoodType
	g.neighborhoodFunc = neighborhood.GetFunc(g.neighborhoodType)
}

// Generation : return the number of generations passed
func (g *Gol) Generation() int {
	return g.generation
}

// NeighborhoodType : return the neighborhood type
func (g *Gol) NeighborhoodType() int {
	return g.neighborhoodType
}

// NeighborhoodTypeString : return the neighborhood type (as string)
func (g *Gol) NeighborhoodTypeString() string {
	return neighborhood.GetName(g.neighborhoodType)
}

// Rows : return the number of rows of the grid
func (g *Gol) Rows() int {
	return g.grid.Rows()
}

// Cols : return the number of columns of the grid
func (g *Gol) Cols() int {
	return g.grid.Cols()
}

// LimitRows : inform if rows are limited or isn't
func (g *Gol) LimitRows() bool {
	return g.grid.LimitRows()
}

// LimitCols : return the number of columns of the grid
func (g *Gol) LimitCols() bool {
	return g.grid.LimitCols()
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
	rows := g.Rows()
	cols := g.Cols()
	nextG := g.Clone().(*Gol)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			aliveNeighborsCount := neighborhood.NeighborsCount(g, i, j, ALIVE, g.neighborhoodFunc)
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
				panic(fmt.Sprintf("Invalid cell %d,%d status", i, j))
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
	clone.InitWithGrid(g.generation, g.neighborhoodType, g.grid.Clone())
	return clone
}
