package output

import (
	"bufio"
	"fmt"
	"os"

	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// SaveToCongolwayFile : prints on stdout the current state of the grid
func (gout *GolOutputer) SaveToCongolwayFile(filename string, fileType string) error {
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
	writer.WriteString(fmt.Sprintf("name: %s\n", gout.name()))
	writer.WriteString(fmt.Sprintf("description: %s\n", gout.description()))
	writer.WriteString(fmt.Sprintf("generation: %d\n", gout.generation()))
	writer.WriteString(fmt.Sprintf("neighborhood_type: %s\n", gout.neighborhoodTypeString()))
	writer.WriteString(fmt.Sprintf("size: %dx%d\n", rows, cols))
	writer.WriteString(fmt.Sprintf("limits: %s\n", gout.limitsString()))
	writer.WriteString("grid_type: dense\n")
	writer.WriteString("grid:\n")

	if fileType == "dense" {
		gout.writeDenseGrid(writer)
	} else if fileType == "dense" {
		gout.writeSparseGrid(writer)
	} else {
		return fmt.Errorf("Invalid file type, expected \"dense\" or \"sparse\", found %s", fileType)
	}

	writer.Flush()
	return nil
}

func (gout *GolOutputer) writeSparseGrid(writer *bufio.Writer) {
	rows := gout.gol.Rows()
	cols := gout.gol.Cols()

	// Count wich status has more cells
	statusCount := make([]int, 0, 2)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			statusCount[gout.get(i, j)]++
		}
	}
	if statusCount[0] >= statusCount[1] {
		writer.WriteString("default: 0\n")
		writer.WriteString("0:\n")
		writer.WriteString(fmt.Sprintf("1: %s\n", gout.coordinateString(0)))
	} else {
		writer.WriteString("default: 1\n")
		writer.WriteString(fmt.Sprintf("0: %s\n", gout.coordinateString(1)))
		writer.WriteString("1:\n")
	}
}

func (gout *GolOutputer) writeDenseGrid(writer *bufio.Writer) {
	rows := gout.gol.Rows()
	cols := gout.gol.Cols()

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if gout.get(i, j) == statuses.ALIVE {
				writer.WriteString("1")
			} else {
				writer.WriteString("0")
			}
		}
		writer.WriteString("\n")
	}
}

func (gout *GolOutputer) coordinateString(value int) string {
	rows := gout.gol.Rows()
	cols := gout.gol.Cols()

	coordinateString := ""
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if gout.get(i, j) == value {
				coordinateString += fmt.Sprintf("(%d,%d)", i, j)
			}
		}
	}
	return coordinateString
}
