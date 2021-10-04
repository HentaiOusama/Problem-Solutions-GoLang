package CustomTools

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Printf("%s\n", p)
	return len(p), nil
}

func GetOutputWriter(size int) *bufio.Writer {
	return bufio.NewWriterSize(new(Writer), size)
}

func ReadLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(strings.TrimSpace(string(str)), "\r\n")
}

func GetInputReader(size int) *bufio.Reader {
	return bufio.NewReaderSize(os.Stdin, size)
}
