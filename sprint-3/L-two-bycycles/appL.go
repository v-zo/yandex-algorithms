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
	days, amounts, bikeCost := readData(reader)

	res1 := findDay(amounts, 0, days-1, bikeCost)
	res2 := findDay(amounts, res1, days-1, 2*bikeCost)

	writer.WriteString(strconv.Itoa(1+res1) + " " + strconv.Itoa(1+res2))
	writer.WriteString("\n")
	writer.Flush()
}

func getEl(s string, n int) int {
	a, _ := strconv.Atoi(string(s[2*n]))
	return a
}

func findDay(amounts string, start int, end int, bikeCost int) int {
	if getEl(amounts, end) < bikeCost {
		return -2
	}

	if end == start {
		return start
	}

	if end-start == 1 {
		if getEl(amounts, start) >= bikeCost {
			return start
		}

		return end
	}

	if getEl(amounts, end-(end-start)/2) >= bikeCost {
		return findDay(amounts, start, end-(end-start)/2, bikeCost)
	} else {
		return findDay(amounts, end-(end-start)/2+1, end, bikeCost)
	}
}

func readData(reader *bufio.Reader) (days int, amounts string, bikeCost int) {
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line3, _ := reader.ReadString('\n')

	days, _ = strconv.Atoi(strings.Trim(line1, "\n"))
	amounts = strings.Trim(line2, "\n")
	bikeCost, _ = strconv.Atoi(line3)

	return
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
