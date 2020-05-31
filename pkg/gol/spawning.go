package gol

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/neighborhood"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// SERIAL : if assigned to Gol.Processes, a serial algorithm (i.e. no concurrency)
// will be used when computing the next generation.
const SERIAL = 1

// CPUS : if assigned to Gol.Processes a number of GO processes equal
// to the number of CPUs of the computer will be used when computing
// the next generation.
const CPUS = -1

// Processes : return the number of GO processes used in
// the computing of the next generation.
// Take account the constants SERIAL and CPUS of this package.
func (g *Gol) Processes() int {
	return g.processes
}

// SetProcesses : set the number of GO processes used in
// the computing of the next generation.
// Take account the constants SERIAL and CPUS of this package.
func (g *Gol) SetProcesses(processes int) {
	g.processes = processes
}

// FastForward : move forward a number of generations
func (g *Gol) FastForward(generations int) base.GolInterface {
	ffg := g.Clone().(*Gol)
	if g.processes == SERIAL {
		for generation := 0; generation < generations; generation++ {
			ffg = ffg.serialNextGeneration().(*Gol)
		}
		return ffg
	}
	for generation := 0; generation < generations; generation++ {
		ffg = ffg.parallelNextGeneration().(*Gol)
	}
	return ffg
}

// ChangeCells : apply cell changes before a new generation
// This method is used by NextGeneration to allow input data
// on the Game of Life instance
func (g *Gol) ChangeCells(changes [][]int) base.GolInterface {
	if changes == nil || len(changes) == 0 {
		return g
	}
	if g.processes == SERIAL {
		return g.serialChangeCells(changes)
	}
	return g.parallelChangeCells(changes)
}

func (g *Gol) serialChangeCells(changes [][]int) base.GolInterface {
	gCopy := g.Clone()
	for _, change := range changes {
		gCopy.Set(change[0], change[1], change[2])
	}
	return gCopy
}

func (g *Gol) parallelChangeCells(changes [][]int) base.GolInterface {
	setRuntimeProcs(g)

	var wg sync.WaitGroup
	wg.Add(len(changes))

	gCopy := g.Clone()
	for _, change := range changes {
		go func(i, j, status int) {
			gCopy.Set(i, j, status)
			wg.Done()
		}(change[0], change[1], change[2])
	}

	wg.Wait()
	return gCopy
}

// NextGeneration : compute the next generation
// If no prior change to the generation of the next game of life
// instance, pass a nil in the place of changes parameter.
func (g *Gol) NextGeneration() base.GolInterface {
	if g.processes == SERIAL {
		return g.serialNextGeneration()
	}
	return g.parallelNextGeneration()
}

// serialNextGeneration : compute the next generation without running threads
func (g *Gol) serialNextGeneration() base.GolInterface {
	rows := g.Rows()
	cols := g.Cols()

	nextG := g.copyWithEmptyGrid().(*Gol)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cellValue := g.nextCell(i, j)
			nextG.Set(i, j, cellValue)
		}
	}
	nextG.generation++
	return nextG
}

// parallelNextGeneration : compute the next generation using threads
func (g *Gol) parallelNextGeneration() base.GolInterface {
	setRuntimeProcs(g)

	nextG := g.copyWithEmptyGrid().(*Gol)

	rows := g.Rows()
	cols := g.Cols()

	var wg sync.WaitGroup
	wg.Add(rows * cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			go func(iIndex int, jIndex int) {
				nextValue := g.nextCell(iIndex, jIndex)
				nextG.Set(iIndex, jIndex, nextValue)
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()
	nextG.generation++
	return nextG
}

func (g *Gol) nextCell(i int, j int) int {
	aliveNeighborsCount := neighborhood.NeighborsCount(g, i, j, statuses.ALIVE, g.neighborhoodFunc)
	// Text from Wikipedia: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
	// Any live cell with two or three live neighbors survives.
	// Any dead cell with three live neighbors becomes a live cell.
	// All other live cells die in the next generation. Similarly, all other dead cells stay dead.
	switch g.Get(i, j) {
	case statuses.ALIVE:
		if g.survivalRule[aliveNeighborsCount] {
			return statuses.ALIVE
		}
		return statuses.DEAD
	case statuses.DEAD:
		if g.birthRule[aliveNeighborsCount] {
			return statuses.ALIVE
		}
		return statuses.DEAD
	default:
		panic(fmt.Sprintf("Invalid cell %d,%d status", i, j))
	}
}

func (g *Gol) copyWithEmptyGrid() base.GolInterface {
	ngGol := new(Gol)
	ngGol.InitWithGrid(g.name, g.description, g.rules, g.generation, g.neighborhoodType, g.grid.CloneEmpty())
	return ngGol
}

func setRuntimeProcs(g base.GolInterface) {
	var processes int
	if g.Processes() == CPUS {
		processes = runtime.NumCPU()
	} else {
		processes = g.Processes()
	}
	runtime.GOMAXPROCS(processes)
}
