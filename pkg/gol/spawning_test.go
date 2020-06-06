package gol

import (
	"testing"

	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

func TestStillLifeNextGeneration(t *testing.T) {
	// Test still-life
	for numOfGenerations := 1; numOfGenerations <= 5; numOfGenerations++ {
		testStillNextGeneration(t, "block.txt", numOfGenerations)
		testStillNextGeneration(t, "bee-hive.txt", numOfGenerations)
		testStillNextGeneration(t, "loaf.txt", numOfGenerations)
		testStillNextGeneration(t, "boat.txt", numOfGenerations)
		testStillNextGeneration(t, "tub.txt", numOfGenerations)
	}
}

func TestOscilatorNextGeneration(t *testing.T) {
	// Test oscilators
	testOscilatorNextGeneration(t, "blinker/gen_0.txt", "blinker/gen_1.txt")
	testOscilatorNextGeneration(t, "beacon/gen_0.txt", "beacon/gen_1.txt")
	testOscilatorNextGeneration(t, "toad/gen_0.txt", "toad/gen_1.txt")
}

func TestBigGridsNextGeneration(t *testing.T) {
	// Test big grids
	testStandardGridNextGeneration(t, "grid1024x1024.txt", "grid1024x1024_gen1.txt", SERIAL, 0)
	testStandardGridNextGeneration(t, "grid1024x1024.txt", "grid1024x1024_gen1.txt", CPUS, 100)
	testStandardGridNextGeneration(t, "grid1024x1024.txt", "grid1024x1024_gen1.txt", CPUS, 200)
	testStandardGridNextGeneration(t, "grid1024x1024.txt", "grid1024x1024_gen1.txt", CPUS, DefaultThreadPoolSize)
	testStandardGridNextGeneration(t, "grid1024x1024.txt", "grid1024x1024_gen1.txt", CPUS, ExplosiveThreadPoolSize)
}

func TestFastForward(t *testing.T) {
	rows := 100
	cols := 100
	randomSeed := int64(42)
	g := NewRandomGol("Random", "", "23/3", "dok", "limited", "limited",
		rows, cols, randomSeed)
	g3 := g.NextGeneration().NextGeneration().NextGeneration().(*Gol)
	ffg3 := g.FastForward(3)

	equalsError := g3.EqualsError(ffg3)
	if equalsError != nil {
		t.Error(equalsError)
	}
}

func TestPriorChanges(t *testing.T) {
	g, gError := readCongolwayFile("still/boat.txt")
	if gError != nil {
		t.Error(gError)
		return
	}

	changes := [][]int{
		{0, 0, statuses.ALIVE},
		{0, 4, statuses.ALIVE},
		{4, 0, statuses.ALIVE},
		{4, 4, statuses.ALIVE},
		{1, 1, statuses.DEAD},
		{1, 2, statuses.DEAD},
	}
	changedG := g.ChangeCells(changes)

	expectedG, expectedGError := readCongolwayFile("expected_changed_boat.txt")
	if expectedGError != nil {
		t.Error(expectedGError)
		return
	}

	equalsError := changedG.EqualsError(expectedG)
	if equalsError != nil {
		t.Error(equalsError)
	}
}

func testStandardGridNextGeneration(t *testing.T, gen0FilePath string, gen1FilePath string, goProcesses int, threadPoolSize int) {
	g0, g0ReadError := readCongolwayFile(gen0FilePath)
	if g0ReadError != nil {
		t.Error(g0ReadError)
	}
	g1, g1ReadError := readCongolwayFile(gen1FilePath)
	if g1ReadError != nil {
		t.Error(g1ReadError)
	}

	g0.SetProcesses(goProcesses)
	if goProcesses != SERIAL {
		g0.SetThreadPoolSize(threadPoolSize)
	}
	actualG1 := g0.NextGeneration().(*Gol)

	if g0.Equals(actualG1) {
		t.Errorf("Standard-life should change after a generation")
	}
	if !g1.Equals(actualG1) {
		t.Errorf("Standard-life should be equal than the expected gol")
	}
}

func testStillNextGeneration(t *testing.T, stillFilePath string, generations int) {
	g0, g0ReadError := readCongolwayFile("still/" + stillFilePath)
	if g0ReadError != nil {
		t.Error(g0ReadError)
	}
	ffg := g0.FastForward(generations)
	if !ffg.GridEquals(g0, "values") {
		t.Errorf("Still-life does not change after a generation")
	}
}

func testOscilatorNextGeneration(t *testing.T, gen0FilePath string, gen1FilePath string) {
	g0, g0ReadError := readCongolwayFile("oscilators/" + gen0FilePath)
	if g0ReadError != nil {
		t.Error(g0ReadError)
	}

	g1, g1ReadError := readCongolwayFile("oscilators/" + gen1FilePath)
	if g1ReadError != nil {
		t.Error(g1ReadError)
	}

	if g0.GridEquals(g1, "values") {
		t.Errorf("Odd oscilator game-of-life generation is wrong. They should be different")
	}

	if !g0.NextGeneration().GridEquals(g1, "values") {
		t.Errorf("Odd oscilator game-of-life generation is wrong. They should be equal (odd generation)")
	}

	if !g1.NextGeneration().GridEquals(g0, "values") {
		t.Errorf("Odd oscilator game-of-life generation is wrong. They should be equal (even generation)")
	}
}
