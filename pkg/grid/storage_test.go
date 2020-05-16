package grid

import (
	"testing"
)

func TestNewStorage(t *testing.T) {
	s := NewStorage(5, 7, 0)
	if s.Rows() != 5 {
		t.Errorf("Invalid rows. Should be %d, found %d", 5, s.Rows())
	}
	if s.Cols() != 7 {
		t.Errorf("Invalid cols. Should be %d, found %d", 7, s.Cols())
	}
	if s.DefaultValue() != 0 {
		t.Errorf("Invalid default value. Should be %d, found %d", 0, s.DefaultValue())
	}
}

func TestStorageGetSet(t *testing.T) {
	s := NewStorage(5, 7, 0)

	expectedValue := 8
	s.Set(1, 2, expectedValue)
	value := s.Get(1, 2)
	if value != expectedValue {
		t.Errorf("Invalid value. Should be %d, found %d", expectedValue, value)
	}
}

func TestStorageClone(t *testing.T) {
	s := NewStorage(5, 7, 0)
	o := s.Clone()
	equalsError := s.EqualsError(o)
	if equalsError != nil {
		t.Error(equalsError)
	}

	if !s.Equals(o) {
		t.Errorf("Should be equal")
	}
}

func TestStorageNotEqual(t *testing.T) {
	s := NewStorage(5, 7, 0)
	s.Set(1, 2, 3)

	// Test that are keys in o1 but not in s
	o1 := s.Clone()
	o1.Set(3, 2, 1)
	o1EqualsError := s.EqualsError(o1)
	if o1EqualsError == nil {
		t.Error("Should be different: o1 and s")
		return
	}
	expectedO1ErrorString := "(3,2) = 1 exists only in the storage grid argument but not in the receiver"
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
	expectedO3ErrorString := "Cells at (1,3) are different: 0 vs 555"
	if o3EqualsError.Error() != expectedO3ErrorString {
		t.Errorf("Expected error: \"%s\". Found: \"%s\"", expectedO3ErrorString, o3EqualsError.Error())
		return
	}
}
