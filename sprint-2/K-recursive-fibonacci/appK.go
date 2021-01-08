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
	n := readData(reader)

	result := fibonacci(n, 1, 0, 0)

	writer.WriteString(strconv.Itoa(result))
	writer.WriteString("\n")

	writer.Flush()
}

func fibonacci(n, p, p0, index int) int {
	if index == n {
		return p
	}

	return fibonacci(n, p+p0, p, index+1)
}

func readData(reader *bufio.Reader) (n int) {
	line1, _ := reader.ReadString('\n')
	n, _ = strconv.Atoi(strings.TrimRight(line1, "\n"))

	return
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}
