package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := openFile("input.txt")
	defer file.close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	process(reader, writer)

	writer.WriteString("\n")
	err := writer.Flush()
	check(err)
}

func process(reader *bufio.Reader, writer io.Writer) {
	yaReader := &YaReader{reader}
	n, seq1, m, seq2 := readData(yaReader)

	res := solve(int16(n), seq1, int16(m), seq2)

	io.WriteString(writer, strconv.Itoa(int(res)))
}

func solve(n int16, seq1 []string, m int16, seq2 []string) (max int) {
	l1 := len(seq1)
	l2 := len(seq2)

	for i := 0; i < l1; i++ {

		for j := 0; j < l2; j++ {

			c := 0
			for i+c < l1 && j+c < l2 && seq1[i+c] == seq2[j+c] {
				c++
			}

			if c > max {
				max = c
			}
		}
	}

	return
}

func readData(reader *YaReader) (n int, seq1 []string, m int, seq2 []string) {
	n = reader.readInt()
	seq1 = reader.readSequence()

	m = reader.readInt()
	seq2 = reader.readSequence()

	return
}

type YaReader struct {
	*bufio.Reader
}

func (reader *YaReader) readString() string {
	line, err := reader.ReadString('\n')
	check(err)
	return strings.TrimRight(line, "\n")
}

func (reader *YaReader) readSequence() []string {
	line := reader.readString()
	trimmedLine := strings.TrimRight(line, "\n")
	return strings.Fields(trimmedLine)
}

func (reader *YaReader) readInt() int {
	line, err := reader.ReadString('\n')
	check(err)
	res, err := strconv.Atoi(strings.TrimRight(line, "\n"))
	check(err)
	return res
}

type File struct {
	*os.File
}

func openFile(path string) *File {
	osFile, err := os.Open(path)
	check(err)

	return &File{osFile}
}

func (file *File) close() {
	err := file.Close()
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
