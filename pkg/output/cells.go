package output

import (
	"bufio"
	"fmt"
	"os"

	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// SaveToCellsFile : prints on stdout the current state of the grid
func (gout *GolOutputer) SaveToCellsFile(filename string) error {
	file, err := os.Create(filename)
	defer file.Close()

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)

	rows := gout.gol.Rows()
	cols := gout.gol.Cols()

	// TODO: write name when name attribute is added to Gol
	writer.WriteString(fmt.Sprintf("!Name: %s\n", gout.name()))
	writer.WriteString(fmt.Sprintf("!%s\n", gout.description()))
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if gout.get(i, j) == statuses.ALIVE {
				writer.WriteString("O")
			} else {
				writer.WriteString(".")
			}
		}
		if i < rows-1 {
			writer.WriteString("\n")
		}
	}

	writer.Flush()
	return nil
}
