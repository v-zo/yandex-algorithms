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
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	yaScanner := &YaScanner{scanner}

	n, m := yaScanner.scanPair()
	adjMap := make(map[int][]int)

	for i := 0; i < n; i++ {
		adjMap[i] = make([]int, n)
		for j := 0; j < n; j++ {
			adjMap[i][j] = 0
		}
	}

	for i := 0; i < m; i++ {
		dot1, dot2 := yaScanner.scanPair()

		adjMap[dot1-1][dot2-1] = 1
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

type YaScanner struct {
	*bufio.Scanner
}

func (scanner *YaScanner) scanPair() (int, int) {
	scanner.Scan()
	line := scanner.Text()
	fields := strings.Fields(line)
	n, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])

	return n, m
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
