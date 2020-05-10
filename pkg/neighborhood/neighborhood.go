package neighborhood

type gridInterface interface {
	Get(i int, j int) int
}

func mooreNeighbors(g gridInterface, i int, j int) []int {
	return []int{
		g.Get(i-1, j-1), g.Get(i-1, j), g.Get(i-1, j+1),
		g.Get(i, j-1), g.Get(i, j+1),
		g.Get(i+1, j-1), g.Get(i+1, j), g.Get(i+1, j+1),
	}
}

func vonNeumanNeighbors(g gridInterface, i int, j int) []int {
	return []int{
		g.Get(i-1, j),
		g.Get(i, j-1), g.Get(i, j+1),
		g.Get(i+1, j),
	}
}

// NeighborsWithStatusCount : number of neighbors alive surronding
// cell i, j of the grid
func NeighborsWithStatusCount(g gridInterface, i int, j int, status int) int {
	aliveCount := 0
	for _, neighborhood := range mooreNeighbors(g, i, j) {
		if neighborhood == status {
			aliveCount++
		}
	}
	return aliveCount
}
