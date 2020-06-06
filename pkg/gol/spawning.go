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

// DefaultThreadPoolSize : default number of threads used in the thread pool
// to compute next generation of a game of life instance.
const DefaultThreadPoolSize = 10

// ExplosiveThreadPoolSize : the thread pool will use a thread for each cell
// do not use it unless you want to experiment with your memory limits.
const ExplosiveThreadPoolSize = -1

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

// ThreadPoolSize : get the number of threads that will be
// used when the parallel next generation algorithm is used.
func (g *Gol) ThreadPoolSize() int {
	return g.threadPoolSize
}

// SetThreadPoolSize : set the number of threads
// that will be used when the parallel next generation algorithm
// is used.
func (g *Gol) SetThreadPoolSize(threadPoolSize int) {
	g.threadPoolSize = threadPoolSize
}

// FastForward : move forward a number of generations
func (g *Gol) FastForward(generations int) base.GolInterface {
	ffg := g.Clone().(*Gol)
	nextGenFunc := g.nextGenerationFunc()
	for generation := 0; generation < generations; generation++ {
		ffg = nextGenFunc(ffg).(*Gol)
	}
	return ffg
}

// NextGeneration : compute the next generation
// If no prior change to the generation of the next game of life
// instance, pass a nil in the place of changes parameter.
func (g *Gol) NextGeneration() base.GolInterface {
	nextGenFunc := g.nextGenerationFunc()
	return nextGenFunc(g)
}

// serialNextGeneration : compute the next generation without running threads
func serialNextGeneration(g *Gol) base.GolInterface {
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

// explosiveParallelNextGeneration : compute the next generation using threads
// creating a thread for each cell. Do not use unless you want to test
// your system's memory limits
func explosiveParallelNextGeneration(g *Gol) base.GolInterface {
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

// parallelNextGeneration : compute the next generation using a thread pool,
func parallelNextGeneration(g *Gol, poolSize int) base.GolInterface {
	setRuntimeProcs(g)

	nextG := g.copyWithEmptyGrid().(*Gol)

	rows := g.Rows()
	cols := g.Cols()

	type pos struct {
		I int
		J int
	}

	jobs := make(chan pos, poolSize)

	var wg sync.WaitGroup

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			wg.Add(1)
			go func(jobs <-chan pos) {
				for pos := range jobs {
					nextValue := g.nextCell(pos.I, pos.J)
					nextG.Set(pos.I, pos.J, nextValue)
				}
				wg.Done()
			}(jobs)
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			jobs <- pos{i, j}
		}
	}
	close(jobs)

	wg.Wait()
	nextG.generation++
	return nextG
}

func (g *Gol) nextGenerationFunc() func(gx *Gol) base.GolInterface {
	if g.processes == SERIAL {
		return serialNextGeneration
	}
	if g.threadPoolSize == ExplosiveThreadPoolSize {
		return explosiveParallelNextGeneration
	}
	return func(gx *Gol) base.GolInterface {
		return parallelNextGeneration(gx, gx.threadPoolSize)
	}
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
