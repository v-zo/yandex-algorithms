package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := openFile("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	Solve(reader, writer)
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	yaReader := &YaReader{reader}
	a, m, s := readData(yaReader)

	output := solve(a, m, s)

	writer.WriteString(output)
	writer.WriteString("\n")
	writer.Flush()
}

func solve(a int, m int, s string) string {
	L := len(s)
	if L == 0 {
		return "0"
	}
	res := int(s[0])
	for i := 1; i < L; i++ {
		res = (a*res)%m + int(s[i])%m
	}

	return strconv.Itoa(res % m)
}

func readData(reader *YaReader) (a int, m int, s string) {
	a = reader.readInt()
	m = reader.readInt()
	s = reader.readString()

	return
}

type YaReader struct {
	*bufio.Reader
}

func (reader *YaReader) readString() string {
	line, _ := reader.ReadString('\n')
	return strings.TrimRight(line, "\n")
}

func (reader *YaReader) readInt() int {
	line, _ := reader.ReadString('\n')
	res, _ := strconv.Atoi(strings.TrimRight(line, "\n"))
	return res
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}
