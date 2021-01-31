package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	q = 1000000007
	m = 1000000009
)

func main() {
	file := openFile("input.txt")
	defer file.close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	processData(reader, writer)

	_, err := writer.WriteString("\n")
	check(err)
	err = writer.Flush()
	check(err)
}

func processData(reader *bufio.Reader, writer io.Writer) {
	yaReader := &YaReader{reader}
	n, k, s := readData(yaReader)

	res := solve(n, k, s)

	_, err := io.WriteString(writer, strings.TrimRight(res, " "))
	check(err)
}

type positionCounter struct {
	pos   int
	count int
}

func powIntMod(N, q, m int) (p int) {
	p = 1

	for i := 0; i < N; i++ {
		p = (p * q) % m
	}

	return
}

func mod(x, m int) int {
	return (x%m + m) % m
}

func solve(n int, k int, s string) (positions string) {
	L := len(s)

	if L == 1 {
		positions = "0"
		return
	}

	indexMap := make(map[int]positionCounter)
	qn := powIntMod(n-1, q, m)

	md := func(x int) int {
		return mod(x, m)
	}

	h := int(s[0])
	for i := 1; i < n; i++ {
		h = (h*q + int(s[i])) % m
	}

	indexMap[h] = positionCounter{0, 1}
	for i := 1; i < L-n+1; i++ {
		h = ((md(h-int(s[i-1])*qn) * q) + int(s[i+n-1])) % m
		iw, has := indexMap[h]
		pos := i
		if has {
			if iw.pos < pos {
				pos = iw.pos
			}
			indexMap[h] = positionCounter{pos, iw.count + 1}
		} else {
			indexMap[h] = positionCounter{pos, 1}
		}
	}

	for _, pc := range indexMap {
		if pc.count >= k {
			num := strconv.Itoa(pc.pos)
			positions += num + " "
		}
	}

	return
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
