package grid

// Interface : minimal Grid interface.
type Interface interface {
	Rows() int
	Cols() int
	LimitRows() bool
	LimitCols() bool
	LimitRowsString() string
	LimitColsString() string
	Get(i int, j int) int
	Set(i int, j int, value int)
	SetAll(value int)
	Equals(other Interface) bool
	EqualsError(other Interface) error
	Clone() Interface
	CloneEmpty() Interface
}
