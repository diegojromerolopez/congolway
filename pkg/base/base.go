package base

type GolInterface interface {
	Init(rows int, cols int, generation int)
	Rows() int
	Cols() int
	Clone() GolInterface
	Get(i int, j int) int
	Set(i int, j int, value int)
	Generation() int
	GridEquals(g GolInterface) bool
	Equals(g GolInterface) bool
	NextGeneration() GolInterface
}
