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
	n, m, north, south := readData(reader)

	output := solve(n, m, north, south)

	writer.WriteString(output)
	writer.WriteString("\n")
	writer.Flush()
}

func solve(n int, m int, north *bufio.Scanner, south *bufio.Scanner) (output string) {
	size := (n + m) / 2

	scanNum := func(scanner *bufio.Scanner) (num int) {
		scanner.Scan()
		num, _ = strconv.Atoi(scanner.Text())

		return
	}

	var merged []int

	a := scanNum(north)
	b := scanNum(south)
	last := 0
	for i := 0; i < size; i++ {
		if a < b {
			merged = append(merged, a)
			a = scanNum(north)
			last = b
		} else {
			merged = append(merged, b)
			b = scanNum(south)
			last = a
		}
	}

	if (n+m)%2 == 0 {
		output = getEvenOutput(last + merged[size-1])
	} else {
		output = strconv.Itoa(last)
	}

	return
}

func getEvenOutput(num int) (output string) {
	s2 := strconv.Itoa(num / 2)
	if num%2 == 1 {
		output = s2 + "." + "5"
	} else {
		output = s2
	}

	return
}

func readData(reader *bufio.Reader) (n int, m int, north *bufio.Scanner, south *bufio.Scanner) {
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line3, _ := reader.ReadString('\n')
	line4, _ := reader.ReadString('\n')

	n, _ = strconv.Atoi(strings.TrimRight(line1, "\n"))
	m, _ = strconv.Atoi(strings.TrimRight(line2, "\n"))

	createScanner := func(line string) *bufio.Scanner {
		//str := strings.TrimRight(line, "\n")
		strReader := strings.NewReader(line)
		strBufReader := bufio.NewReader(strReader)
		scanner := bufio.NewScanner(strBufReader)
		scanner.Split(bufio.ScanWords)

		return scanner
	}

	north = createScanner(line3)
	south = createScanner(line4)

	return
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
