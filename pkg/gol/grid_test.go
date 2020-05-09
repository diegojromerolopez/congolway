package gol

import (
	"testing"
)

func TestCountAliveNeighborhood(t *testing.T) {
	g, gReadError := readGolFromTextFile("3x3.txt")
	if gReadError != nil {
		t.Error(gReadError)
	}
	grid := g.grid
	expectedAliveNeighborhoodSize := [][]int{
		{2, 3, 3},
		{3, 3, 2},
		{1, 2, 2},
	}
	for i := 0; i < grid.rows; i++ {
		for j := 0; j < grid.cols; j++ {
			if expectedAliveNeighborhoodSize[i][j] != grid.aliveNeighborsCount(i, j) {
				t.Errorf(
					"Cell %d,%d should have %d neighbors but it has %d",
					i, j,
					expectedAliveNeighborhoodSize[i][j],
					grid.aliveNeighborsCount(i, j),
				)
			}
		}
	}

}
