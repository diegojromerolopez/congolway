package output

import (
	"fmt"
	"strings"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// GolOutputer : tasked with writting in several devices the passed
// Game of Life (gol.Gol object)
type GolOutputer struct {
	//TODO: use base.GolInterface instead
	gol *gol.Gol
}

// NewGolOutputer : returns a new pointer to GolOutputer
func NewGolOutputer(g *gol.Gol) *GolOutputer {
	return &GolOutputer{g}
}

// SaveToFile : prints on stdout the current state of the grid
func (gout *GolOutputer) SaveToFile(filename string) error {
	lastDotIndex := strings.LastIndex(filename, ".")
	if lastDotIndex < 0 {
		return fmt.Errorf("File \"%s\" has no extension. Only .txt and .cells files are allowed", filename)
	}
	fileExtension := filename[lastDotIndex:]

	if fileExtension == ".txt" {
		return gout.SaveToCongolwayFile(filename, "dense")
	} else if fileExtension == ".cells" {
		return gout.SaveToCellsFile(filename)
	}
	return fmt.Errorf("File extension \"%s\" not recognized. Only .txt and .cells are allowed", fileExtension)
}

func (gout *GolOutputer) name() string {
	return gout.gol.Name()
}

func (gout *GolOutputer) description() string {
	return gout.gol.Description()
}

func (gout *GolOutputer) generation() int {
	return gout.gol.Generation()
}

func (gout *GolOutputer) rules() string {
	return gout.gol.Rules()
}

func (gout *GolOutputer) get(i int, j int) int {
	return gout.gol.Get(i, j)
}

func (gout *GolOutputer) neighborhoodTypeString() string {
	return gout.gol.NeighborhoodTypeString()
}

func (gout *GolOutputer) limitsString() string {
	limitsStr := ""
	if gout.gol.LimitRows() && gout.gol.LimitCols() {
		limitsStr += "rows, cols"
	} else if gout.gol.LimitRows() {
		limitsStr += "rows"
	} else if gout.gol.LimitCols() {
		limitsStr += "cols"
	} else {
		panic("Impossible condition")
	}
	return limitsStr
}

// Stdout : prints on stdout the current state of the grid
// Return the numbers of lines used
func (gout *GolOutputer) Stdout(cellStringCorresp map[int]string) int {
	if cellStringCorresp == nil {
		cellStringCorresp = map[int]string{
			statuses.DEAD:  "░",
			statuses.ALIVE: "█",
		}
	}
	g := gout.gol
	rows := g.Rows()
	cols := g.Cols()
	fmt.Printf("%s\n", g.Name())
	fmt.Printf("Generation: %d\n", g.Generation())
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cellString := cellStringCorresp[g.Get(i, j)]
			fmt.Print(cellString)
		}
		fmt.Print("\n")
	}
	return rows + 2
}
