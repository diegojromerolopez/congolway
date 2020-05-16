package grid

import (
	"fmt"
	"sync"
)

type _Key struct {
	i int
	j int
}

// Storage : grid storage as a sparse matrix
type Storage struct {
	cells        *sync.Map
	defaultValue int
	rows         int
	cols         int
}

// NewStorage : create a new grid storage (i.e. a dictionary-key-based sparse matrix)
func NewStorage(rows, cols, defaultValue int) *Storage {
	return &Storage{
		rows:         rows,
		cols:         cols,
		defaultValue: defaultValue,
		cells:        new(sync.Map),
	}
}

// Rows : return the number of rows of the grid storage
func (s *Storage) Rows() int {
	return s.rows
}

// Cols : return the number of columns of the grid storage
func (s *Storage) Cols() int {
	return s.cols
}

// DefaultValue : return the default value of the grid storage
// The default value is that one that is returned when cell (i,j)
// is not found
func (s *Storage) DefaultValue() int {
	return s.defaultValue
}

// Get : get the value of the cell (ALICE, DEAD)
//	in the i, j coordinates
func (s *Storage) Get(i, j int) int {
	s.assertIndexes(i, j)
	value, valueExists := s.cells.Load(_Key{i, j})
	if valueExists {
		return value.(int)
	}
	return s.defaultValue
}

// Set : set the value of the cell in the i, j coordinates
func (s *Storage) Set(i, j, value int) {
	s.assertIndexes(i, j)
	s.cells.Store(_Key{i, j}, value)
}

// SetAll : set a value to all ceels
func (s *Storage) SetAll(value int) {
	s.defaultValue = value
	s.cells = new(sync.Map)
}

// Equals : inform if two grids have the same cell value
// for each position.s
func (s *Storage) Equals(o *Storage) bool {
	return s.EqualsError(o) == nil
}

// EqualsError : inform if two grids have the same dimensions and
// the same cell values for each position.
func (s *Storage) EqualsError(other *Storage) error {
	if s.rows != other.rows {
		return fmt.Errorf("Rows are different: %d vs %d", s.rows, other.rows)
	}
	if s.cols != other.cols {
		return fmt.Errorf("Cols are different: %d vs %d", s.cols, other.cols)
	}

	// Check that every key in s has the same value in other
	var myKey _Key
	var myValue int
	var otherValue int
	var otherKeyExists bool
	cellsAreEqual := true
	s.cells.Range(func(key, value interface{}) bool {
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
		_, sKeyExists := s.cells.Load(key)
		if !sKeyExists {
			otherHasNonCommonItems = true
			nonCommonOtherKey = key.(_Key)
			nonCommonOtherValue = value.(int)
			return false
		}
		return true
	})

	if otherHasNonCommonItems {
		return fmt.Errorf("(%d,%d) = %d exists only in the storage grid argument but not in the receiver",
			nonCommonOtherKey.i, nonCommonOtherKey.j, nonCommonOtherValue)
	}

	// Otherwise, everything went well, no errors
	return nil
}

// Clone : clone the grid storage in a new grid storage
func (s *Storage) Clone() *Storage {
	clone := NewStorage(s.rows, s.cols, s.defaultValue)
	s.cells.Range(func(key, value interface{}) bool {
		clone.cells.Store(key, value)
		return true
	})
	return clone
}

func (s *Storage) assertIndexes(i, j int) {
	if i < 0 || i >= s.rows {
		panic(fmt.Sprintf("Invalid row index: %d not in [0, %d]", i, s.rows-1))
	}
	if j < 0 || j >= s.cols {
		panic(fmt.Sprintf("Invalid col index: %d not in [0, %d]", j, s.cols-1))
	}
}
