package gol

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/grid"
	"github.com/diegojromerolopez/congolway/pkg/neighborhood"
)

// Gol : game of life
type Gol struct {
	name             string
	description      string
	grid             *grid.Grid
	generation       int
	neighborhoodType int
	neighborhoodFunc neighborhood.Func
	rules            string
	survivalRule     map[int]bool // Poor's man set
	birthRule        map[int]bool // Poor's man set
	processes        int
}

// NewGol : creates a game of life
func NewGol(name, description, rules, gridType, rowsLimitation, colsLimitation string, rows, cols, generation int) *Gol {
	g := new(Gol)
	gr := grid.NewGrid(rows, cols, rowsLimitation, colsLimitation, gridType)
	g.InitWithGrid(name, description, rules, generation, neighborhood.MOORE, gr)
	return g
}

// NewRandomGol : creates a new random game of life
func NewRandomGol(name, description, rules, gridType, rowsLimitation, colsLimitation string, rows, cols int, randomSeed int64) *Gol {
	g := new(Gol)
	gr := grid.NewRandomGrid(rows, cols, rowsLimitation, colsLimitation, gridType, randomSeed)
	g.InitWithGrid(name, description, rules, 0, neighborhood.MOORE, gr)
	return g
}

// InitFromConf : initialize a Game of Life instance
func (g *Gol) InitFromConf(name, description string, rows, cols int, gconf *base.GolConf) {
	g.init(name, description,
		gconf.Rules(), gconf.GridType(),
		gconf.RowLimitation(), gconf.ColLimitation(),
		rows, cols, gconf.Generation(), gconf.NeighborhoodType())
}

// InitWithGrid : initialize a Game of Life instance
func (g *Gol) InitWithGrid(name, description, rules string, generation, neighborhoodType int, gr *grid.Grid) {
	g.name = name
	g.description = description
	g.generation = generation
	g.SetRules(rules)
	g.neighborhoodType = neighborhoodType
	g.neighborhoodFunc = neighborhood.GetFunc(g.neighborhoodType)
	g.grid = gr
	g.processes = CPUS
}

// Name : return the name of this Game of life instance
func (g *Gol) Name() string {
	return g.name
}

// Description : return the description of this Game of life instance
func (g *Gol) Description() string {
	return g.description
}

// Rules : return the rules of the game of life
// as a string with the survival/birth format.
// See https://www.conwaylife.com/wiki/Life_1.05
func (g *Gol) Rules() string {
	return g.rules
}

// SetRules : set rules according with the
// survival/birth format.
// See https://www.conwaylife.com/wiki/Life_1.05
func (g *Gol) SetRules(rules string) {
	g.rules = rules
	rulesParts := strings.Split(rules, "/")
	// TODO: put common code in function
	g.survivalRule = make(map[int]bool)
	survivalRule := rulesParts[0]
	for i := 0; i < len(survivalRule); i++ {
		sr, srErr := strconv.Atoi(survivalRule[i : i+1])
		if srErr != nil {
			panic(srErr.Error())
		}
		g.survivalRule[sr] = true
	}
	g.birthRule = make(map[int]bool)
	birthRule := rulesParts[1]
	for i := 0; i < len(birthRule); i++ {
		br, brErr := strconv.Atoi(birthRule[i : i+1])
		if brErr != nil {
			panic(brErr.Error())
		}
		g.birthRule[br] = true
	}
}

// Generation : return the number of generations passed
func (g *Gol) Generation() int {
	return g.generation
}

// SetGeneration : set generation passed
func (g *Gol) SetGeneration(generation int) {
	g.generation = generation
}

// NeighborhoodType : return the neighborhood type
func (g *Gol) NeighborhoodType() int {
	return g.neighborhoodType
}

// NeighborhoodTypeString : return the neighborhood type (as string)
func (g *Gol) NeighborhoodTypeString() string {
	return neighborhood.StringFromType(g.neighborhoodType)
}

// SetNeighborhoodType : return the neighborhood type
func (g *Gol) SetNeighborhoodType(neighborhoodType int) {
	neighborhood.AssertType(neighborhoodType)
	g.neighborhoodType = neighborhoodType
}

// SetNeighborhoodTypeString : return the neighborhood type (as string)
func (g *Gol) SetNeighborhoodTypeString(neighborhoodType string) {
	g.neighborhoodType = neighborhood.TypeFromString(neighborhoodType)
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

// SetLimitRows : set if rows are limited or isn't
func (g *Gol) SetLimitRows(limitRows bool) {
	g.grid.SetLimitRows(limitRows)
}

// LimitCols : return the number of columns of the grid
func (g *Gol) LimitCols() bool {
	return g.grid.LimitCols()
}

// SetLimitCols : set if cols are limited or isn't
func (g *Gol) SetLimitCols(limitRows bool) {
	g.grid.SetLimitCols(limitRows)
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

// SetAll : set the value to all cells
func (g *Gol) SetAll(value int) {
	g.grid.SetAll(value)
}

// Equals : inform if two game of life instances have the same data
func (g *Gol) Equals(o base.GolInterface) bool {
	other := o.(*Gol)
	simpleAttributesAreEqual := g.name == other.name &&
		g.description == other.description &&
		g.generation == other.generation &&
		g.rules == other.rules &&
		g.neighborhoodType == other.neighborhoodType &&
		g.processes == other.processes

	return simpleAttributesAreEqual && g.grid.Equals(other.grid, "values")
}

// GridEquals : inform if two game of life instances have the same data,
//	ignoring the difference in generations value
func (g *Gol) GridEquals(o base.GolInterface, mode string) bool {
	other := o.(*Gol)
	return g.grid.Equals(other.grid, mode)
}

// EqualsError : inform if two game of life instances have the same data
// by returning an error if different, or nil otherwise.
func (g *Gol) EqualsError(o base.GolInterface) error {
	other := o.(*Gol)

	if g.name != other.name {
		return fmt.Errorf("Names are different: \"%s\" vs \"%s\"", g.name, other.name)
	}

	if g.description != other.description {
		return fmt.Errorf("Descriptions are different: \"%s\" vs \"%s\"", g.description, other.description)
	}

	if g.rules != other.rules {
		return fmt.Errorf("Rules are different: %s vs %s", g.rules, other.rules)
	}

	if g.generation != other.generation {
		return fmt.Errorf("Generations are different: %d vs %d", g.generation, other.generation)
	}

	if g.neighborhoodType != other.neighborhoodType {
		return fmt.Errorf("Neighborhoodtype are different: %s vs %s", g.NeighborhoodTypeString(), other.NeighborhoodTypeString())
	}

	if g.processes != other.processes {
		return fmt.Errorf("Processes are different: %d vs %d", g.processes, other.processes)
	}

	return g.grid.EqualsError(other.grid, "values")
}

// Clone : clone a game of life instance
func (g *Gol) Clone() base.GolInterface {
	clone := new(Gol)
	clone.InitWithGrid(g.name, g.description, g.rules, g.generation, g.neighborhoodType, g.grid.Clone())
	clone.SetProcesses(g.processes)
	return clone
}

// DbgStdout : show a matrix to ease debugging
func (g *Gol) DbgStdout() {
	rows := g.Rows()
	cols := g.Cols()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(g.Get(i, j))
		}
		fmt.Print("\n")
	}
}

func (g *Gol) init(name, description, rules, gridType, rowsLimitation, colsLimitation string, rows, cols, generation, neighborhoodType int) {
	g.name = name
	g.description = description
	g.grid = grid.NewGrid(rows, cols, rowsLimitation, colsLimitation, gridType)
	g.SetRules(rules)
	g.generation = generation
	g.neighborhoodType = neighborhoodType
	g.neighborhoodFunc = neighborhood.GetFunc(g.neighborhoodType)
	g.processes = CPUS
}
