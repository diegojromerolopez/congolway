package base

// GolInterface : minimal Gol interface.
type GolInterface interface {
	Init(rows int, cols int,
		rowsLimitation string, colsLimitation string,
		generation int, neighborhoodType int)
	Rows() int
	Cols() int
	LimitRows() bool
	LimitCols() bool
	Clone() GolInterface
	Get(i int, j int) int
	Set(i int, j int, value int)
	Generation() int
	GridEquals(g GolInterface) bool
	Equals(g GolInterface) bool
	NeighborhoodTypeString() string
	GetProcesses() int
	SetProcesses(processes int)
	NextGeneration() GolInterface
}
