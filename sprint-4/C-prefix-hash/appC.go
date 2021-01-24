package main

import (
	"bufio"
	"os"
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

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	yaReader := &YaReader{reader}
	a, m, s := readData(yaReader)

	solve(a, m, s, reader, writer)

	writer.Flush()
}

func solve(a int, m int, s string, reader *bufio.Reader, writer *bufio.Writer) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())

	hashStr := createHashStr(a, m, s)

	for i := 0; i < M; i++ {
		scanner.Scan()
		pair := strings.Fields(scanner.Text())
		lo, _ := strconv.Atoi(pair[0])
		hi, _ := strconv.Atoi(pair[1])
		writer.WriteString(hashStr(lo-1, hi-1))
		writer.WriteString("\n")
	}
}

func horner(a int, m int, s string) []int {
	L := len(s)
	h := make([]int, L)
	h[0] = int(s[0])
	for i := 1; i < L; i++ {
		h[i] = (h[i-1]*a + int(s[i])) % m
	}

	return h
}

func createHashStr(a int, m int, s string) func(int, int) string {
	L := len(s)
	pows := powIntMod(a, m, L)
	h := horner(a, m, s)
	aModM := a % m

	return func(lo int, hi int) string {
		delta := hi - lo

		if delta < lo {
			return strconv.Itoa(horner(a, m, s[lo:hi+1])[delta])
		}

		res := h[delta]
		pd := pows[delta]
		for i := 0; i < lo; i++ {
			ss := res - (pd * int(s[i]))
			res = mod(ss, m)*aModM + int(s[i+delta+1])%m
		}

		return strconv.Itoa(res % m)
	}
}

func mod(x, m int) int {
	return (x%m + m) % m
}

func powIntMod(a, m, L int) (p []int) {
	p = make([]int, L)
	p[0] = 1

	for i := 1; i < L; i++ {
		p[i] = p[i-1] * a % m
	}

	return
}

func readData(reader *YaReader) (a int, m int, s string) {
	a = reader.readInt()
	m = reader.readInt()
	s = reader.readString()

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
