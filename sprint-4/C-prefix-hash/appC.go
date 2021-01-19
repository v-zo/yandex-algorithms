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
		writer.WriteString(hashStr(lo, hi))
		writer.WriteString("\n")
	}
}

func createHashStr(a int, m int, s string) func(int, int) string {
	L := len(s)
	//var memo []int
	memo := []int{int(s[0])}
	for i := 1; i < L; i++ {
		memo = append(memo, (a*memo[i-1])%m+int(s[i])%m)
	}

	return func(lo int, hi int) string {
		if L == 0 {
			return "0"
		}

		if lo == hi {
			return strconv.Itoa(int(s[lo-1]) % m)
		}

		//res := int(s[lo-1])
		//for i := lo; i < hi; i++ {
		//	res = memo[lo-1] + int(s[i])%m
		//}

		res := memo[hi-1] - memo[lo-1]

		return strconv.Itoa(res % m)
	}
}

//func hashStr(a int, m int, s string, lo int, hi int) string {
//	if  len(s) == 0 {
//		return "0"
//	}
//
//	if lo==hi {
//		return strconv.Itoa(int(s[lo-1])%m)
//	}
//
//	res := int(s[lo-1])
//	for i := lo; i < hi; i++ {
//		res = (a*res)%m + int(s[i])%m
//	}
//
//	return strconv.Itoa(res % m)
//}

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
