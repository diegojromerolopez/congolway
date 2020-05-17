package animator

import (
	"fmt"
	"time"

	"github.com/diegojromerolopez/congolway/pkg/gol"
	"github.com/diegojromerolopez/congolway/pkg/output"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// MakeStdout : make a terminal animation for some generations
func MakeStdout(g *gol.Gol, generations int, delay int) error {
	delayInDuration, delayInDurationError := time.ParseDuration(fmt.Sprintf("%dms", delay))
	if delayInDurationError != nil {
		return delayInDurationError
	}
	cellStringCorrespondence := map[int]string{
		statuses.DEAD:  "░",
		statuses.ALIVE: "█",
	}
	for generationI := 0; generationI < generations; generationI++ {
		gout := output.NewGolOutputer(g)
		terminalRowsUsed := gout.Stdout(cellStringCorrespondence)
		time.Sleep(delayInDuration)
		fmt.Printf("\033[%dA", terminalRowsUsed)
		g = g.NextGeneration().(*gol.Gol)
	}
	return nil
}
