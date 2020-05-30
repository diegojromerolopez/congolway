package grid

import (
	"fmt"
	"sync"
)

type _Key struct {
	i int
	j int
}

// Dok : dictionary of keys grid (sparse matrix)
type Dok struct {
	cells        *sync.Map
	defaultValue int
	rows         int
	cols         int
}

// NewDok : create a new grid dictOfKeys (i.e. a dictionary-key-based sparse matrix)
func NewDok(rows int, cols int, defaultValue int) *Dok {
	dok := new(Dok)
	dok.cells = new(sync.Map)
	dok.rows = rows
	dok.cols = cols
	return dok
}

// Rows : return the number of rows of the grid dictOfKeys
func (dok *Dok) Rows() int {
	return dok.rows
}

// Cols : return the number of columns of the grid dictOfKeys
func (dok *Dok) Cols() int {
	return dok.cols
}

// DefaultValue : return the default value of the grid dictOfKeys
// The default value is that one that is returned when cell (i,j)
// is not found
func (dok *Dok) DefaultValue() int {
	return dok.defaultValue
}

// Get : get the value of the cell (ALICE, DEAD)
//	in the i, j coordinates
func (dok *Dok) Get(i, j int) int {
	dok.assertIndexes(i, j)
	value, valueExists := dok.cells.Load(_Key{i, j})
	if valueExists {
		return value.(int)
	}
	return dok.defaultValue
}

// Set : set the value of the cell in the i, j coordinates
func (dok *Dok) Set(i, j, value int) {
	dok.assertIndexes(i, j)
	if value == dok.defaultValue {
		dok.cells.Delete(_Key{i, j})
	} else {
		dok.cells.Store(_Key{i, j}, value)
	}
}

// SetAll : set a value to all cells
func (dok *Dok) SetAll(value int) {
	dok.defaultValue = value
	dok.cells = new(sync.Map)
}

// Equals : inform if two grids have the same cell value
// for each position.s
func (dok *Dok) Equals(o CellsStorer) bool {
	return dok.EqualsError(o) == nil
}

// EqualsError : inform if two grids have the same dimensions and
// the same cell values for each position.
func (dok *Dok) EqualsError(o CellsStorer) error {
	dimensionsError := dok.equalDimensionsError(o)
	if dimensionsError != nil {
		return dimensionsError
	}

	other := o.(*Dok)

	// Check that every key in s has the same value in other
	var myKey _Key
	var myValue int
	var otherValue int
	var otherKeyExists bool
	cellsAreEqual := true
	dok.cells.Range(func(key, value interface{}) bool {
		oValue, oKeyExists := other.cells.Load(key)
		otherKeyExists = oKeyExists
		myKey = key.(_Key)
		myValue := value.(int)
		if !otherKeyExists {
			cellsAreEqual = false
			return false
		}
		otherValue = oValue.(int)
		if myValue != otherValue {
			otherKeyExists = true
			cellsAreEqual = false
			return false
		}
		return true
	})

	if !cellsAreEqual {
		if otherKeyExists {
			return fmt.Errorf("Cells at (%d,%d) are different: %d vs %d",
				myKey.i, myKey.j, myValue, otherValue)
		}
		return fmt.Errorf("Cells at (%d,%d) are different: %d vs NOVALUE",
			myKey.i, myKey.j, myValue)
	}

	// Check that other has no more keys than s
	otherHasNonCommonItems := false
	var nonCommonOtherKey _Key
	var nonCommonOtherValue int
	other.cells.Range(func(key, value interface{}) bool {
		_, sKeyExists := dok.cells.Load(key)
		if !sKeyExists {
			otherHasNonCommonItems = true
			nonCommonOtherKey = key.(_Key)
			nonCommonOtherValue = value.(int)
			return false
		}
		return true
	})

	if otherHasNonCommonItems {
		return fmt.Errorf("(%d,%d) = %d exists only in the dok argument but not in the receiver",
			nonCommonOtherKey.i, nonCommonOtherKey.j, nonCommonOtherValue)
	}

	// Otherwise, everything went well, no errors
	return nil
}

// EqualValues : check value by value if both
// cell storers have the same values
func (dok *Dok) EqualValues(o CellsStorer) bool {
	return dok.EqualValuesError(o) != nil
}

// EqualValuesError : check value by value if both
// cell storers have the same values. Return an error
// if that's not the case
func (dok *Dok) EqualValuesError(o CellsStorer) error {
	dimensionsError := dok.equalDimensionsError(o)
	if dimensionsError != nil {
		return dimensionsError
	}

	for i := 0; i < dok.rows; i++ {
		for j := 0; j < dok.cols; j++ {
			if dok.Get(i, j) != o.Get(i, j) {
				return fmt.Errorf("Cells at (%d,%d) are different: %d vs %d",
					i, j, dok.Get(i, j), o.Get(i, j))
			}
		}
	}
	return nil
}

// Clone : clone the grid dictOfKeys in a new grid dictOfKeys
func (dok *Dok) Clone() CellsStorer {
	clone := NewDok(dok.rows, dok.cols, dok.defaultValue)
	dok.cells.Range(func(key, value interface{}) bool {
		clone.cells.Store(key, value)
		return true
	})
	return clone
}

// CloneEmpty : create a new grid with the same size but empty
func (dok *Dok) CloneEmpty() CellsStorer {
	return NewDok(dok.rows, dok.cols, dok.defaultValue)
}

func (dok *Dok) assertIndexes(i, j int) {
	if i < 0 || i >= dok.rows {
		panic(fmt.Sprintf("Invalid row index: %d not in [0, %d]", i, dok.rows-1))
	}
	if j < 0 || j >= dok.cols {
		panic(fmt.Sprintf("Invalid col index: %d not in [0, %d]", j, dok.cols-1))
	}
}

func (dok *Dok) equalDimensionsError(o CellsStorer) error {
	oRows := o.Rows()
	if dok.rows != oRows {
		return fmt.Errorf("Rows are different: %d vs %d", dok.rows, oRows)
	}
	oCols := o.Cols()
	if dok.cols != oCols {
		return fmt.Errorf("Cols are different: %d vs %d", dok.cols, oCols)
	}
	return nil
}
