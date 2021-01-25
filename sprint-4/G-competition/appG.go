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

	Solve(reader, writer)

	err := writer.Flush()
	check(err)
}

func Solve(reader *bufio.Reader, writer io.Writer) {
	yaReader := &YaReader{reader}
	n, sequence := readData(yaReader)

	res := sol(n, sequence)

	io.WriteString(writer, strconv.Itoa(res))
}

type Value struct {
	val    int
	hasKey bool
}

type Map map[int]Value

func (m Map) put(key, val int) {
	m[key] = Value{val, true}
}

func (m Map) get(key int) int {
	return m[key].val
}

func (m Map) has(key int) bool {
	return m[key].hasKey
}

func sol(n int, sequence []int) (res int) {
	if n == 0 {
		return
	}

	if n == 1 {
		return
	}

	h := make(Map)
	h.put(0, -1)
	count := 0
	for i := 0; i < n; i++ {
		count += sequence[i]
		if h.has(count) {
			res = max(res, i-h.get(count))
		} else {
			h.put(count, i)
		}
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func readData(reader *YaReader) (n int, sequence []int) {
	n = reader.readInt()
	if n == 0 {
		sequence = []int{}
		return
	}

	line := reader.readString()

	arr := strings.Split(line, " ")
	for i := 0; i < n; i++ {
		num := 1
		if arr[i] == "0" {
			num = -1
		}
		sequence = append(sequence, num)
	}

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
