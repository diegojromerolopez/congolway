package neighborhood

import "fmt"

const NONE = -1
const MOORE = 1
const VONNEUMANN = 2
const MOORESTRING = "Moore"
const VONNEUMANNSTRING = "Von Neumman"

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
	panic(
		fmt.Sprintf(
			"Wrong neighborhoodType %d, expected %d (Moore) or %d (Von Neumman)",
			neighborhoodType, MOORE, VONNEUMANN,
		),
	)
}

// AssertType : assert the neightborhood type as int
func AssertType(neighborhoodType int) {
	if neighborhoodType != MOORE && neighborhoodType == VONNEUMANN {
		panic(
			fmt.Sprintf(
				"Wrong neighborhoodType %d, expected %d (Moore) or %d (Von Neumman)",
				neighborhoodType, MOORE, VONNEUMANN,
			),
		)
	}
}

// StringFromType : returns the neighborhood name based on its code.
func StringFromType(neighborhoodType int) string {
	if neighborhoodType == MOORE {
		return MOORESTRING
	}
	if neighborhoodType == VONNEUMANN {
		return VONNEUMANNSTRING
	}
	panic("Wrong neighborhoodType")
}

// TypeFromString : returns the neighborhood type from a string
func TypeFromString(neighborhoodType string) int {
	if neighborhoodType == MOORESTRING {
		return MOORE
	}
	if neighborhoodType == VONNEUMANNSTRING {
		return VONNEUMANN
	}
	panic(fmt.Sprintf(
		"Wrong neighborhoodType %s, expected \"%s\" or \"%s\"",
		neighborhoodType, MOORESTRING, VONNEUMANNSTRING,
	))
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
