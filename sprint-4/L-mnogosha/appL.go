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

func solve(n int, k int, s string) (positions string) {
	L := len(s)

	if L == 1 {
		positions = "0"
		return
	}

	indexMap := make(map[string]positionCounter)

	for i := 0; i < L-n+1; i++ {
		word := s[i : i+n]
		if _, has := indexMap[word]; has {
			pos := i
			if indexMap[word].pos < pos {
				pos = indexMap[word].pos
			}
			indexMap[word] = positionCounter{pos, indexMap[word].count + 1}
		} else {
			indexMap[word] = positionCounter{i, 1}
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
