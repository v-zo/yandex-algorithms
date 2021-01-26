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

	res := solve(n, seq1, m, seq2)

	io.WriteString(writer, strconv.Itoa(res))
}

func solve(n int, seq1 []string, m int, seq2 []string) (max int) {
	L := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		L[i] = make(map[int]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if seq1[i] == seq2[j] {
				if i == 0 || j == 0 {
					L[i][j] = 1
				} else {
					L[i][j] = L[i-1][j-1] + 1
				}
				if L[i][j] > max {
					max = L[i][j]
				}
			} else {
				L[i][j] = 0
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
