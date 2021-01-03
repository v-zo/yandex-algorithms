package main

import (
	"bufio"
	"os"
	"sort"
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

type Sortable []string

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	_, arr := readData(reader)

	sort.Sort(arr)

	writer.WriteString(strings.Join(arr, ""))
	writer.WriteString("\n")

	writer.Flush()
}

func compare(a string, b string) bool {
	if len(a) > len(b) {
		delta := len(a) - len(b)
		bNew := b + a[:delta]

		return a > bNew
	}

	if len(a) < len(b) {
		delta := len(b) - len(a)
		aNew := a + b[:delta]

		return aNew > b
	}

	return a > b
}

func (s Sortable) Len() int      { return len(s) }
func (s Sortable) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Sortable) Less(i, j int) bool {
	a, b := s[i], s[j]

	return compare(a, b)
}

func readData(reader *bufio.Reader) (int, Sortable) {
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')

	n, _ := strconv.Atoi(line1)
	arr := strings.Fields(line2)

	return n, arr
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
