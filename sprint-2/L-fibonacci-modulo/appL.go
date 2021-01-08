package main

import (
	"bufio"
	"math"
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
	n, k := readData(reader)

	result := fibonacciModulo(n, k, 1, 0)

	writer.WriteString(strconv.Itoa(result))
	writer.WriteString("\n")

	writer.Flush()
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func fibonacciModulo(n, k, p, p0 int) int {
	for i := 0; i < n; i++ {
		p, p0 = p0, p
		p = (p + p0) % powInt(10, k)
	}

	return p
}

func readData(reader *bufio.Reader) (n, k int) {
	line1, _ := reader.ReadString('\n')
	fields := strings.Fields(strings.TrimRight(line1, "\n"))
	n, _ = strconv.Atoi(fields[0])
	k, _ = strconv.Atoi(fields[1])

	return
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}
