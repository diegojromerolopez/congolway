package input

import (
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/base"
)

func assertGolIsRight(t *testing.T, filename string, name string, description string,
	rows int, cols int, limitRows bool, limitCols bool,
	generation int, expectedCells [][]int, g base.GolInterface) {

	if g.Name() != name {
		t.Errorf("Loaded name is wrong, got: %s, must be: %s.", g.Name(), name)
		return
	}

	if g.Description() != description {
		t.Errorf("Loaded description is wrong, got: %s, must be: %s.", g.Description(), description)
		return
	}

	if g.Generation() != generation {
		t.Errorf("Loaded generation is wrong, got: %d, must be: %d.", g.Generation(), generation)
		return
	}

	if g.NeighborhoodTypeString() != "Moore" {
		t.Errorf("Loaded neighborhood is wrong, got: %s, must be: Moore.", g.NeighborhoodTypeString())
		return
	}

	if g.Rows() != rows {
		t.Errorf("Loaded number of rows is wrong, got: %d, must be: %d.", g.Rows(), rows)
		return
	}
	if g.Cols() != cols {
		t.Errorf("Loaded number of cols is wrong, got: %d, must be: %d.", g.Cols(), cols)
		return
	}

	if g.LimitRows() != limitRows {
		if limitRows {
			t.Errorf("Should limit rows, but it isn't.")
		} else {
			t.Errorf("Shouldn't limit rows, but it is.")
		}
		return
	}
	if g.LimitCols() != limitCols {
		if limitCols {
			t.Errorf("Should limit cols, but it isn't.")
		} else {
			t.Errorf("Shouldn't limit cols, but it is.")
		}
		return
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			gIJ := g.Get(i, j)
			expectedIJ := expectedCells[i][j]
			if gIJ != expectedIJ {
				t.Errorf("Invalid cell at %d,%d. It got: %d, must be: %d.", i, j, gIJ, expectedIJ)
			}
		}
	}
}
