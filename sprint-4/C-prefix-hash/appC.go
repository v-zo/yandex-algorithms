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

	solve(a, m, s, reader, writer)

	writer.Flush()
}

func solve(a int, m int, s string, reader *bufio.Reader, writer *bufio.Writer) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < M; i++ {
		scanner.Scan()
		pair := strings.Fields(scanner.Text())
		lo, _ := strconv.Atoi(pair[0])
		hi, _ := strconv.Atoi(pair[1])
		writer.WriteString(hashStr(a, m, s, lo, hi))
		writer.WriteString("\n")
	}
}

func hashStr(a int, m int, s string, lo int, hi int) string {
	if len(s) == 0 {
		return "0"
	}

	if lo == hi {
		return strconv.Itoa(int(s[lo-1]) % m)
	}

	res := int(s[lo-1])
	for i := lo; i < hi; i++ {
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
