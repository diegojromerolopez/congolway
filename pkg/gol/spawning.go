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

// GetProcesses : return the number of GO processes used in
// the computing of the next generation.
// Take account the constants SERIAL and CPUS of this package.
func (g *Gol) GetProcesses() int {
	return g.processes
}

// SetProcesses : set the number of GO processes used in
// the computing of the next generation.
// Take account the constants SERIAL and CPUS of this package.
func (g *Gol) SetProcesses(processes int) {
	g.processes = processes
}

// NextGeneration : compute the next generation
func (g *Gol) NextGeneration() base.GolInterface {
	if g.processes == SERIAL {
		return g.serialNextGeneration()
	}
	return g.parallelNextGeneration()
}

// SerialNextGeneration : compute the next generation
func (g *Gol) serialNextGeneration() base.GolInterface {
	rows := g.Rows()
	cols := g.Cols()

	nextG := g.createNextGenerationGol().(*Gol)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cellValue := g.nextCell(i, j)
			nextG.Set(i, j, cellValue)
		}
	}
	nextG.generation++
	return nextG
}

func (g *Gol) parallelNextGeneration() base.GolInterface {
	if g.processes == CPUS {
		g.processes = runtime.NumCPU()
	}
	runtime.GOMAXPROCS(g.processes)

	nextG := g.createNextGenerationGol().(*Gol)

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
		if aliveNeighborsCount == 2 || aliveNeighborsCount == 3 {
			return statuses.ALIVE
		}
		return statuses.DEAD
	case statuses.DEAD:
		if aliveNeighborsCount == 3 {
			return statuses.ALIVE
		}
		return statuses.DEAD
	default:
		panic(fmt.Sprintf("Invalid cell %d,%d status", i, j))
	}
}

func (g *Gol) createNextGenerationGol() base.GolInterface {
	ngGol := new(Gol)
	ngGol.InitWithGrid(g.generation, g.neighborhoodType, g.grid.CloneEmpty())
	return ngGol
}
