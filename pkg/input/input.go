package input

import (
	"github.com/diegojromerolopez/congolway/pkg/base"
)

// GolReader : tasked with reading a Game of Life from files
type GolReader struct {
	readGol base.GolInterface
}

// NewGolReader : returns a new pointer to GolReader
func NewGolReader(g base.GolInterface) *GolReader {
	return &GolReader{g}
}
