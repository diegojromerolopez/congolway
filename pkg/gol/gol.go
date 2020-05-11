package gol

import (
	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/grid"
	"github.com/diegojromerolopez/congolway/pkg/neighborhood"
)

// Gol : game of life
type Gol struct {
	grid             *grid.Grid
	generation       int
	neighborhoodType int
	neighborhoodFunc neighborhood.Func
	processes        int
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
