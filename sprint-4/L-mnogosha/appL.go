package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	q = 1000
	m = 1000000009
)

func main() {
	file := openFile("input.txt")
	defer file.close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	process(reader, writer)

	_, err := writer.WriteString("\n")
	check(err)
	err = writer.Flush()
	check(err)
}

func process(reader *bufio.Reader, writer io.Writer) {
	yaReader := &YaReader{reader}
	n, k, s := readData(yaReader)

	res := solve(n, k, s)

	_, err := io.WriteString(writer, res)
	check(err)
}

func prefixHash(s string, n int) int {
	h := int(s[0])

	if len(s) == 1 {
		return h
	}

	for i := 1; i < n; i++ {
		h = (h*q + int(s[i])) % m
	}

	return h % m
}

func solve(n int, k int, s string) string {
	qn := powIntMod(n - 1)
	md := func(x int) int {
		return mod(x, m)
	}

	positions := make(map[int][]int)
	hash := prefixHash(s, n)
	positions[hash] = []int{0}

	for i := 0; i < len(s)-n; i++ {
		hash = md(md(md(hash)-int(s[i])*qn)*q + int(s[n+i]))
		if positions[hash] == nil {
			positions[hash] = []int{i + 1}
		} else {
			positions[hash] = append(positions[hash], i+1)
		}
	}

	var arr []string
	for _, pos := range positions {
		if len(pos) >= k {
			arr = append(arr, strconv.Itoa(pos[0]))
		}
	}

	return strings.Join(arr, " ")
}

func mod(x, m int) int {
	return (x%m + m) % m
}

func powIntMod(n int) int {
	p := 1
	for i := 0; i < n; i++ {
		p = p * q % m
	}

	return p
}

func readData(reader *YaReader) (n int, k int, s string) {
	seq := reader.readSequenceInt()
	n, k = seq[0], seq[1]
	s = reader.readString()

	return
}

type YaReader struct {
	*bufio.Reader
}

func (reader *YaReader) readString() string {
	line, err := reader.ReadString('\n')
	if err == io.EOF {
		return line
	}
	check(err)
	return strings.TrimRight(line, "\n")
}

func (reader *YaReader) readSequence() []string {
	line := reader.readString()
	trimmedLine := strings.TrimRight(line, "\n")
	return strings.Fields(trimmedLine)
}

func (reader *YaReader) readSequenceInt() (seqInt []int) {
	seqStr := reader.readSequence()
	seqInt = make([]int, len(seqStr))
	var err error
	for i, s := range seqStr {
		seqInt[i], err = strconv.Atoi(s)
		check(err)
	}

	return
}

func (reader *YaReader) readInt() int {
	line, err := reader.ReadString('\n')
	if err != io.EOF {
		line = strings.TrimRight(line, "\n")
	}
	check(err)
	res, err := strconv.Atoi(line)
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
