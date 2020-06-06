package gol

import (
	"sync"

	"github.com/diegojromerolopez/congolway/pkg/base"
)

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
