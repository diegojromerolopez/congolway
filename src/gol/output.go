package gol

import (
	"bufio"
	"fmt"
	"os"
)

// Stdout : prints on stdout the current state of the grid
func (g *Gol) Stdout() {
	grid := g.grid
	for i := 0; i < grid.rows; i++ {
		for j := 0; j < grid.cols; j++ {
			if grid.get(i, j) == ALIVE {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

// SaveToFile : prints on stdout the current state of the grid
func (g *Gol) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	defer file.Close()

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	grid := g.grid

	writer.WriteString("CONGOLWAY\n")
	writer.WriteString("version: 1\n")
	writer.WriteString(fmt.Sprintf("generation: %d\n", g.generation))
	writer.WriteString(fmt.Sprintf("size: %dx%d\n", grid.rows, grid.cols))
	writer.WriteString("grid:\n")

	for i := 0; i < grid.rows; i++ {
		for j := 0; j < grid.cols; j++ {
			if grid.get(i, j) == ALIVE {
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
