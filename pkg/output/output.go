package output

import (
	"bufio"
	"fmt"
	"os"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// GolOutputer : tasked with writting in several devices the passed
// Game of Life (gol.Gol object)
type GolOutputer struct {
	gol *gol.Gol
}

// NewGolOutputer : returns a new pointer to GolOutputer
func NewGolOutputer(g *gol.Gol) *GolOutputer {
	return &GolOutputer{g}
}

func (gout *GolOutputer) get(i int, j int) int {
	return gout.gol.Get(i, j)
}

func (gout *GolOutputer) generation() int {
	return gout.gol.Generation()
}

func (gout *GolOutputer) neighborhoodTypeString() string {
	return gout.gol.NeighborhoodTypeString()
}

// Stdout : prints on stdout the current state of the grid
func (gout *GolOutputer) Stdout() {
	rows := gout.gol.Rows()
	cols := gout.gol.Cols()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if gout.get(i, j) == statuses.ALIVE {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

// SaveToFile : prints on stdout the current state of the grid
func (gout *GolOutputer) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	defer file.Close()

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)

	rows := gout.gol.Rows()
	cols := gout.gol.Cols()

	writer.WriteString("CONGOLWAY\n")
	writer.WriteString("version: 1\n")
	writer.WriteString(fmt.Sprintf("generation: %d\n", gout.generation()))
	writer.WriteString(fmt.Sprintf("neighborhood_type: %s\n", gout.neighborhoodTypeString()))
	writer.WriteString(fmt.Sprintf("size: %dx%d\n", rows, cols))
	writer.WriteString("grid:\n")

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if gout.get(i, j) == statuses.ALIVE {
				writer.WriteString("X")
			} else {
				writer.WriteString(" ")
			}
		}
		writer.WriteString("\n")
	}
	writer.Flush()
	return nil
}
