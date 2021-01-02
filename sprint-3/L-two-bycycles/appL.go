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
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line3, _ := reader.ReadString('\n')

	days, _ := strconv.Atoi(strings.Trim(line1, "\n"))
	amountStrings := strings.Split(strings.Trim(line2, "\n"), " ")
	bikeCost, _ := strconv.Atoi(line3)

	var amounts []int
	for _, str := range amountStrings {
		a, _ := strconv.Atoi(str)
		amounts = append(amounts, a)
	}

	res1, res2 := findTwoDays(amounts, days, bikeCost)

	r1 := conv(res1)
	r2 := conv(res2)

	writer.WriteString(r1 + " " + r2)
	writer.WriteString("\n")
	writer.Flush()
}

func conv(i int) string {
	if i == -1 {
		return "-1"
	}

	return strconv.Itoa(i + 1)
}

func findTwoDays(amounts []int, days int, bikeCost int) (res1 int, res2 int) {
	res2 = -1
	res1 = -1

	if days == 1 {
		if amounts[0] >= bikeCost {
			res1 = 0
		}
		if amounts[0] >= 2*bikeCost {
			res2 = 0
		}

		return
	}

	res1 = findDay(amounts, 0, days-1, bikeCost)

	if res1 > -1 {
		res2 = findDay(amounts, res1+1, days-1, 2*bikeCost)
	}

	return
}

func findDay(amounts []int, start int, end int, bikeCost int) int {
	if end == start {
		if amounts[start] >= bikeCost {
			return start
		}

		return -1
	}

	if end-start == 1 {
		if amounts[start] >= bikeCost {
			return start
		}

		if amounts[start+1] >= bikeCost {
			return start + 1
		}

		return -1
	}

	if end-start == 2 {
		if amounts[start] >= bikeCost {
			return start
		}

		if amounts[start+1] >= bikeCost {
			return start + 1
		}

		if amounts[end] >= bikeCost {
			return end
		}

		return -1
	}

	var endNew int
	var startNew int

	if amounts[end/2] >= bikeCost {
		endNew = end / 2
		startNew = start
	} else {
		endNew = end
		startNew = end/2 + 1
	}

	return findDay(amounts, startNew, endNew, bikeCost)
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
