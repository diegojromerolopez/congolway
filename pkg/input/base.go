package input

import (
	"bufio"
	"io"
	"os"
)

type fileReader struct {
	currentLine *string
	file        *os.File
	reader      *bufio.Reader
}

// CurrentLine : return a pointer to the current line
func (fr *fileReader) CurrentLine() *string {
	return fr.currentLine
}

func (fr *fileReader) ReadLine() error {
	line, err := fr.reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			fr.currentLine = &line
		} else {
			fr.currentLine = nil
		}
		return err
	}
	lineWithoutEndline := line[:len(line)-1]
	fr.currentLine = &lineWithoutEndline
	return nil
}

func (fr *fileReader) SeekStart() {
	fr.file.Seek(0, io.SeekStart)
	fr.reader = bufio.NewReader(fr.file)
	fr.currentLine = nil
}

func newFileReader(file *os.File) *fileReader {
	return &fileReader{nil, file, bufio.NewReader(file)}
}
