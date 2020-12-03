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
	if n%2 == 0 {
		writer.WriteString("2")
		writer.WriteString(" ")
		FindFactors(n/2, writer)

		return
	}
	for i := 3; i <= n; i += 2 {
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
