package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file := openFile("input.txt")
	n := readData(file)
	file.Close()

	writer := bufio.NewWriter(os.Stdout)

	FindFactors(n, writer)

	writer.WriteString("\n")
	writer.Flush()
}

func FindFactors(n int, writer *bufio.Writer) {
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			writer.WriteString(strconv.Itoa(i))
			writer.WriteString(" ")
			FindFactors(n/i, writer)

			return
		}
	}
}

func readData(inputFile *os.File) int {
	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	return n
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
