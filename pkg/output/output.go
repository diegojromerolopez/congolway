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

func (g *GolOutputer) get(i int, j int) int {
	return g.gol.Get(i, j)
}

func (g *GolOutputer) generation() int {
	return g.gol.Generation()
}

// Stdout : prints on stdout the current state of the grid
func (g *GolOutputer) Stdout() {
	rows := g.gol.Rows()
	cols := g.gol.Cols()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if g.get(i, j) == statuses.ALIVE {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

// SaveToFile : prints on stdout the current state of the grid
func (g *GolOutputer) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	defer file.Close()

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)

	rows := g.gol.Rows()
	cols := g.gol.Cols()

	writer.WriteString("CONGOLWAY\n")
	writer.WriteString("version: 1\n")
	writer.WriteString(fmt.Sprintf("generation: %d\n", g.generation()))
	writer.WriteString(fmt.Sprintf("size: %dx%d\n", rows, cols))
	writer.WriteString("grid:\n")

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if g.get(i, j) == statuses.ALIVE {
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
