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

	scanner := bufio.NewScanner(yaReader)
	scanner.Split(bufio.ScanLines)

	readIntPair := func() (int, int) {
		scanner.Scan()
		line := scanner.Text()
		fields := strings.Fields(line)
		n, _ := strconv.Atoi(fields[0])
		m, _ := strconv.Atoi(fields[1])

		return n, m
	}

	n, m := readIntPair()
	adjMap := make(map[int][]int)

	for i := 0; i < n; i++ {
		adjMap[i] = []int{0}
	}

	for i := 0; i < m; i++ {
		dot1Plus, dot2 := readIntPair()
		dot1 := dot1Plus - 1

		adjMap[dot1][0] = adjMap[dot1][0] + 1
		adjMap[dot1] = append(adjMap[dot1], dot2)
	}

	for i := 0; i < n; i++ {
		outLine := SplitToString(adjMap[i], " ")
		io.WriteString(writer, outLine)
	}
}

func SplitToString(a []int, sep string) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}

	res := strings.Join(b, sep)
	return res + "\n"
}

type YaReader struct {
	*bufio.Reader
}

//
//func (reader *YaReader) readString() string {
//	line, err := reader.ReadString('\n')
//	check(err)
//	return strings.TrimRight(line, "\n")
//}
//
//func (reader *YaReader) readPair() (int, int) {
//	line, err := reader.ReadString('\n')
//	check(err)
//	arr := strings.Split(line, " ")
//	n, err := strconv.Atoi(strings.TrimRight(arr[0], "\n"))
//	check(err)
//	m, err := strconv.Atoi(strings.TrimRight(arr[1], "\n"))
//	check(err)
//
//	return n, m
//}

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
