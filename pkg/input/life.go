package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/diegojromerolopez/congolway/pkg/base"
)

// ReadLifeFile : create a new Game of life from a .life file.
// See the following links for more information:
// - 1.05: https://www.conwaylife.com/wiki/Life_1.05
// - 1.06: https://www.conwaylife.com/wiki/Life_1.06
func (gr *GolReader) ReadLifeFile(filename string, gconf *base.GolConf) (base.GolInterface, error) {
	file, fileError := os.Open(filename)
	defer file.Close()

	if fileError != nil {
		return nil, fileError
	}

	reader := bufio.NewReader(file)
	headerLine, headerLineError := reader.ReadString('\n')
	if headerLineError != nil {
		return nil, headerLineError
	}
	headerText := strings.TrimSuffix(headerLine, "\n")
	if headerText == "#Life 1.05" {
		return gr.ReadLife105File(filename, gconf)
	}
	if headerText == "#Life 1.06" {
		return gr.ReadLife106File(filename, gconf)
	}
	return nil, fmt.Errorf("Invalid header \"%s\" for a Life file", headerText)
}
