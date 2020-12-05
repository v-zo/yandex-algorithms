package main

import (
	"bufio"
	"os"
)

type ProblemSolver struct {
	scanner     *bufio.Scanner
	writer      *bufio.Writer
	homeCounter int
}

func main() {
	file := openFile("input.txt")
	defer file.Close()

	a, b := readNumbers(file)

	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(a)
	writer.WriteString("\n")
	writer.WriteString(b)
	writer.WriteString("\n")

	writer.Flush()
}

func readNumbers(input *os.File) (string, string) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	a := scanner.Text()
	scanner.Scan()
	b := scanner.Text()

	return a, b
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
