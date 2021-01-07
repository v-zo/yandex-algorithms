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
	res2 := -2
	if res1 != -2 {
		res2 = findDay(amounts, res1, days-1, 2*bikeCost)
	}

	writer.WriteString(strconv.Itoa(1+res1) + " " + strconv.Itoa(1+res2))
	writer.WriteString("\n")
	writer.Flush()
}

func getEl(s []string, n int) int {
	a, _ := strconv.Atoi(s[n])
	return a
}

func findDay(amounts []string, start int, end int, bikeCost int) int {
	if getEl(amounts, end) < bikeCost {
		return -2
	}

	if getEl(amounts, start) >= bikeCost {
		return start
	}

	if end == start || end-start == 1 {
		return end
	}

	mid := end - (end-start)/2
	if getEl(amounts, mid) >= bikeCost {
		return findDay(amounts, start, mid, bikeCost)
	} else {
		return findDay(amounts, mid+1, end, bikeCost)
	}
}

func readData(reader *bufio.Reader) (days int, amounts []string, bikeCost int) {
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line3, _ := reader.ReadString('\n')

	days, _ = strconv.Atoi(strings.Trim(line1, "\n"))
	amounts = strings.Split(strings.Trim(line2, "\n"), " ")
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
