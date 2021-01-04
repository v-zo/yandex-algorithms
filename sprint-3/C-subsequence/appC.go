package main

import (
	"bufio"
	"os"
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
	s, t := readData(reader)

	k := 0
	for i := 0; i < len(t) && k < len(s); i++ {
		if t[i] == s[k] {
			k++
		}
	}

	var result string

	if k == len(s) {
		result = "True"
	} else {
		result = "False"
	}

	writer.WriteString(result)
	writer.WriteString("\n")

	writer.Flush()
}

func readData(reader *bufio.Reader) (s string, t string) {
	first, _ := reader.ReadString('\n')
	t, _ = reader.ReadString('\n')

	s = strings.Trim(first, "\n")

	return
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
