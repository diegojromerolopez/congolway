package base

// GolInterface : minimal Gol interface.
type GolInterface interface {
	Init(name string, description string, generation int, rows int, cols int,
		rowsLimitation string, colsLimitation string, rules string, neighborhoodType int)
	Name() string
	Description() string
	Rows() int
	Cols() int
	LimitRows() bool
	LimitCols() bool
	Clone() GolInterface
	Get(i int, j int) int
	Set(i int, j int, value int)
	SetAll(value int)
	Rules() string
	DbgStdout()
	Generation() int
	GridEquals(g GolInterface) bool
	Equals(g GolInterface) bool
	EqualsError(g GolInterface) error
	NeighborhoodTypeString() string
	GetProcesses() int
	SetProcesses(processes int)
	NextGeneration() GolInterface
}
