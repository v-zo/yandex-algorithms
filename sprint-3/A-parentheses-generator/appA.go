package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file := openFile("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	Solve(reader, writer)
}

type Solver struct {
	ans []string
	N   int
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	n := readData(reader)

	s := &Solver{nil, n}

	s.backtrack("", 0, 0)

	for _, v := range s.ans {
		writer.WriteString(v)
		writer.WriteString("\n")
	}

	writer.Flush()
}

func (s *Solver) backtrack(S string, left int, right int) {
	if len(S) == 2*s.N {
		s.ans = append(s.ans, S)
		return
	}

	if left < s.N {
		s.backtrack(S+"(", left+1, right)
	}

	if right < left {
		s.backtrack(S+")", left, right+1)
	}
}

func readData(reader *bufio.Reader) int {
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(line)

	return n
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
