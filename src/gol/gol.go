package gol

// Gol : game of life
type Gol struct {
	grid       *Grid
	generation int
}

// NewGol : creates a game of life
func NewGol(rows int, cols int) *Gol {
	g := new(Gol)
	g.grid = NewGrid(rows, cols)
	g.generation = 0
	return g
}

// NewRandomGol : creates a new random game of life
func NewRandomGol(rows int, cols int, randomSeed int64) *Gol {
	g := new(Gol)
	g.grid = NewRandomGrid(rows, cols, randomSeed)
	g.generation = 0
	return g
}

// NextGeneration : compute the next generation
func (g *Gol) NextGeneration() {
	grid := g.grid
	for i := 0; i < grid.rows; i++ {
		for j := 0; j < grid.cols; j++ {
			aliveNeighborsCount := grid.countAliveNeighbors(i, j)
			// Text from Wikipedia: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
			// Any live cell with two or three live neighbors survives.
			// Any dead cell with three live neighbors becomes a live cell.
			// All other live cells die in the next generation. Similarly, all other dead cells stay dead.
			if grid.get(i, j) == ALIVE {
				if aliveNeighborsCount > 1 {
					grid.set(i, j, ALIVE)
				} else {
					grid.set(i, j, DEAD)
				}
			} else {
				if aliveNeighborsCount > 3 {
					grid.set(i, j, ALIVE)
				}
			}
		}
	}
	g.generation++
}

// Equals : inform if two game of life instances have the same data
func (g *Gol) Equals(other *Gol) bool {
	return g.grid.equals(other.grid) && g.generation == other.generation
}

// Clone : clone a game of life instance
func (g *Gol) Clone() *Gol {
	clone := new(Gol)
	clone.generation = g.generation
	clone.grid = g.grid.clone()
	return clone
}
