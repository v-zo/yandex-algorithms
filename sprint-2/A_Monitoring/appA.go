package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file := openFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lm := initM(scanner)
	writer := bufio.NewWriter(os.Stdout)

	value, done := lm.next()

	counter := 0
	for !done {
		if lm.m-1 == counter {
			writer.WriteString(value)
			writer.WriteString("\n")
			counter = 0
		} else {
			writer.WriteString(value + " ")
			counter++
		}

		value, done = lm.next()
	}

	writer.Flush()
}

type LinkedMatrix struct {
	i     int
	j     int
	n     int
	m     int
	input [][]string
}

func initM(scanner *bufio.Scanner) *LinkedMatrix {
	lm := LinkedMatrix{}
	lm.i = 0
	lm.j = 0

	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	lm.n = n
	lm.m = m

	var inputM = make([][]string, m)
	for j := 0; j < m; j++ {
		inputM[j] = make([]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			cur := scanner.Text()
			inputM[j][i] = cur
		}

	}

	lm.input = inputM

	return &lm
}

func (lm *LinkedMatrix) next() (value string, done bool) {
	if lm.n == lm.j {
		return "0", true
	}

	value = lm.input[lm.i][lm.j]

	if lm.m-lm.i == 1 {
		lm.i = 0
		lm.j = lm.j + 1
	} else {
		lm.i = lm.i + 1
	}

	return value, false
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
