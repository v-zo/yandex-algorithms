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

	solveProblem(reader, writer)
}

func solveProblem(reader *bufio.Reader, writer *bufio.Writer) {
	yaReader := &YaReader{reader}
	_, s := readData(yaReader)

	output := solve(s)

	for _, ot := range output {
		writer.WriteString(ot)
		writer.WriteString("\n")
	}

	writer.Flush()
}

func solve(words []string) (result []string) {
	hashes := make(map[int32][]int)

	for i, w := range words {
		hashes[hash(w)] = append(hashes[hash(w)], i)
	}

	for _, indexes := range hashes {
		var entry []string
		for _, index := range indexes {
			entry = append(entry, strconv.Itoa(index))
		}

		result = append(result, strings.Join(entry, " "))
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return
}

func hash(s string) (r int32) {
	for _, ch := range s {
		r += ch
	}

	r += int32(len(s))

	return
}

func readData(reader *YaReader) (n int, s []string) {
	n = reader.readInt()
	s = strings.Fields(reader.readString())

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
