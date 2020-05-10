package neighborhood

const NONE = -1
const MOORE = 1
const VONNEUMANN = 2

type gettable interface {
	Get(i int, j int) int
}

// Func : neighborhood function type
type Func func(g gettable, i int, j int) []int

func mooreNeighbors(g gettable, i int, j int) []int {
	return []int{
		g.Get(i-1, j-1), g.Get(i-1, j), g.Get(i-1, j+1),
		g.Get(i, j-1), g.Get(i, j+1),
		g.Get(i+1, j-1), g.Get(i+1, j), g.Get(i+1, j+1),
	}
}

func vonNeumannNeighbors(g gettable, i int, j int) []int {
	return []int{
		g.Get(i-1, j),
		g.Get(i, j-1), g.Get(i, j+1),
		g.Get(i+1, j),
	}
}

// GetFunc : returns the neighborhood function based on its code.
func GetFunc(neighborhoodType int) Func {
	if neighborhoodType == MOORE {
		return mooreNeighbors
	}
	if neighborhoodType == VONNEUMANN {
		return vonNeumannNeighbors
	}
	panic("Wrong neighborhoodType")
}

// GetName : returns the neighborhood name based on its code.
func GetName(neighborhoodType int) string {
	if neighborhoodType == MOORE {
		return "Moore"
	}
	if neighborhoodType == VONNEUMANN {
		return "Von Neumann"
	}
	panic("Wrong neighborhoodType")
}

// NeighborsCount : number of neighbors alive surronding
// cell i, j of the grid
func NeighborsCount(g gettable, i int, j int, status int, neighborhoodFunc Func) int {
	neighborsCount := 0
	for _, neighborhood := range neighborhoodFunc(g, i, j) {
		if neighborhood == status {
			neighborsCount++
		}
	}
	return neighborsCount
}
