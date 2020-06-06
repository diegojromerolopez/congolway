package grid

import (
	"reflect"
	"testing"
)

func TestCellsStorerFactory(t *testing.T) {
	defer func() {
		errorString := recover().(string)

		if errorString != "Invalid grid type: whatever. Only \"dense\" or \"dok\" are accepted as gridType values" {
			t.Errorf("Wrong panic message: %s", errorString)
		}
	}()

	denseGrid := CellsStorerFactory(10, 20, "dense")
	denseGridType := reflect.TypeOf(denseGrid)
	if denseGridType.String() != "*grid.Dense" {
		t.Errorf("Expecting *grid.Dense struct, found %s", denseGridType.String())
	}

	dokGrid := CellsStorerFactory(100, 200, "dok")
	dokGridType := reflect.TypeOf(dokGrid)
	if dokGridType.String() != "*grid.Dok" {
		t.Errorf("Expecting *grid.Dok struct, found %s", dokGridType.String())
	}

	CellsStorerFactory(100, 200, "whatever")
}
