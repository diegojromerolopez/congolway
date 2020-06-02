package output

import (
	"bufio"
	"fmt"
	"os"

	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// SaveToLifeFile : save the game of life instance to a
// life (.life) file. An extra parameter is used to
// choose between 1.05 and 1.06 version.
// See the following links for more information:
// - 1.05: https://www.conwaylife.com/wiki/Life_1.05
// - 1.06: https://www.conwaylife.com/wiki/Life_1.06
func (gout *GolOutputer) SaveToLifeFile(filename string, version string) error {
	if version == "1.05" {
		return gout.SaveToLife105File(filename)
	}
	if version == "1.06" {
		return gout.SaveToLife106File(filename)
	}
	return fmt.Errorf("Version %s for Life file format is not recognized. "+
		"Only 1.05 and 1.06 versions are available", version)
}

// SaveToLife105File : save the game of life instance to a
// 1.05 .life file (1.05: https://www.conwaylife.com/wiki/Life_1.05)
func (gout *GolOutputer) SaveToLife105File(filename string) error {
	return fmt.Errorf("Life 1.05 file format is not supported yet")
}

// SaveToLife106File : save the game of life instance to a
// 1.06 .life file (1.06: https://www.conwaylife.com/wiki/Life_1.06)
func (gout *GolOutputer) SaveToLife106File(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	writer.WriteString("#Life 1.06")
	rows := gout.gol.Rows()
	cols := gout.gol.Cols()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if gout.gol.Get(i, j) == statuses.ALIVE {
				writer.WriteString(fmt.Sprintf("\n%d %d", i, j))
			}
		}
	}
	writer.Flush()
	return nil
}
