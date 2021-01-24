package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := openFile("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)

	Solve(reader)
}

func Solve(reader *bufio.Reader) {
	yaReader := &YaReader{reader}
	s := readData(yaReader)

	res := sol(s)

	fmt.Println(res)
}

func sol(s string) int {
	m := 0
	maxLength := 0
	charIndex := make(map[uint8]int, 256)
	for i := uint8(0); i < uint8(255); i++ {
		charIndex[i] = -1
	}

	for i := 0; i < len(s); i++ {
		m = max(charIndex[s[i]]+1, m)
		charIndex[s[i]] = i
		maxLength = max(maxLength, i-m+1)
	}

	return maxLength
}

func max(a int, b int) int {
	if b > a {
		return b
	}

	return a
}

func readData(reader *YaReader) string {
	return reader.readString()
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
