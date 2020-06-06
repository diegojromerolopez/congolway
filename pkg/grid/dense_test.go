package grid

import (
	"testing"
)

func TestNewDense(t *testing.T) {
	s := NewDense(5, 7)
	if s.Rows() != 5 {
		t.Errorf("Invalid rows. Should be %d, found %d", 5, s.Rows())
	}
	if s.Cols() != 7 {
		t.Errorf("Invalid cols. Should be %d, found %d", 7, s.Cols())
	}
}

func TestDenseGetSet(t *testing.T) {
	s := NewDense(5, 7)

	expectedValue := 8
	s.Set(1, 2, expectedValue)
	value := s.Get(1, 2)
	if value != expectedValue {
		t.Errorf("Invalid value. Should be %d, found %d", expectedValue, value)
	}
}

func TestDenseSetAll(t *testing.T) {
	s := NewDense(5, 7)

	expectedValue := 8
	s.SetAll(expectedValue)
	for i := 0; i < s.Rows(); i++ {
		for j := 0; j < s.Cols(); j++ {
			if s.Get(i, j) != expectedValue {
				t.Errorf("Invalid value found at %d,%d. Should be %d, found %d", i, j, expectedValue, s.Get(i, j))
			}
		}
	}
}

func TestDenseClone(t *testing.T) {
	s := NewDense(5, 7)
	s.Set(0, 0, 100)
	s.Set(1, 1, 100)
	s.Set(2, 2, 100)
	s.Set(3, 3, 100)
	s.Set(4, 4, 100)
	o := s.Clone()
	equalsError := s.EqualsError(o)
	if equalsError != nil {
		t.Error(equalsError)
		return
	}

	equalValuesError := s.EqualValuesError(o)
	if equalValuesError != nil {
		t.Error(equalValuesError)
		return
	}

	if !s.EqualValues(o) {
		t.Errorf("Values should be equal")
	}

	if !s.Equals(o) {
		t.Errorf("Should be equal")
	}
}

func TestDenseCloneEmpty(t *testing.T) {
	s := NewDense(5, 7)
	s.Set(0, 0, 100)
	s.Set(1, 1, 100)
	s.Set(2, 2, 100)
	s.Set(3, 3, 100)
	s.Set(4, 4, 100)
	o := s.CloneEmpty()

	for i := 0; i < 5; i++ {
		if s.Get(i, i) == o.Get(i, i) {
			t.Errorf("Values %d,%d should be different", i, i)
		}
		if o.Get(i, i) != 0 {
			t.Errorf("Value %d,%d should be 0", i, i)
		}
	}

	if s.Rows() != o.Rows() {
		t.Errorf("Rows should be equal")
	}

	if s.Cols() != o.Cols() {
		t.Errorf("Cols should be equal")
	}

	if s.EqualValues(o) {
		t.Errorf("Values should be different")
	}

	if s.Equals(o) {
		t.Errorf("Should be different")
	}
}

func TestDenseNotEqual(t *testing.T) {
	s := NewDense(5, 7)
	s.Set(1, 2, 3)

	// Test that are keys in o1 but not in s
	o1 := s.Clone()
	o1.Set(3, 2, 1)
	o1EqualsError := s.EqualsError(o1)
	if o1EqualsError == nil {
		t.Error("Should be different: o1 and s")
		return
	}
	expectedO1ErrorString := "Cells at (3,2) are different: 0 vs 1"
	if o1EqualsError.Error() != expectedO1ErrorString {
		t.Errorf("Expected error: \"%s\". Found: \"%s\"", expectedO1ErrorString, o1EqualsError.Error())
		return
	}

	// Test that are keys in s but not in o2
	o2 := s.Clone()
	s.Set(1, 3, 3)
	o2EqualsError := s.EqualsError(o2)
	if o2EqualsError == nil {
		t.Error("Should be different: o2 and s")
		return
	}

	// Test that are keys in s and in o3 but with different values
	o3 := s.Clone()
	o3.Set(1, 3, 555)
	o3EqualsError := s.EqualsError(o3)
	if o3EqualsError == nil {
		t.Error("Should be different: o3 and s")
		return
	}
	expectedO3ErrorString := "Cells at (1,3) are different: 3 vs 555"
	if o3EqualsError.Error() != expectedO3ErrorString {
		t.Errorf("Expected error: \"%s\". Found: \"%s\"", expectedO3ErrorString, o3EqualsError.Error())
		return
	}
}

func TestDenseNotEqualRows(t *testing.T) {
	s1 := NewDense(5, 7)
	s2 := NewDense(8, 3)
	expectedErrorString := "Rows are different: 5 vs 8"

	equalsError := s1.EqualsError(s2)

	if equalsError == nil {
		t.Error("Should be different: s1 and s2 have different number of rows")
		return
	}

	if equalsError.Error() != expectedErrorString {
		t.Errorf("Should have returned: \"%s\" but returned \"%s\"", expectedErrorString, equalsError.Error())
		return
	}
}

func TestDenseNotEqualCols(t *testing.T) {
	s1 := NewDense(5, 7)
	s2 := NewDense(5, 3)
	expectedErrorString := "Cols are different: 7 vs 3"

	equalsError := s1.EqualsError(s2)

	if equalsError == nil {
		t.Error("Should be different: s1 and s2 have different number of cols")
		return
	}

	if equalsError.Error() != expectedErrorString {
		t.Errorf("Should have returned: \"%s\" but returned \"%s\"", expectedErrorString, equalsError.Error())
		return
	}
}
