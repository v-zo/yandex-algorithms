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
	amounts := strings.Split(strings.Trim(line2, "\n"), " ")
	bikeCost, _ := strconv.Atoi(line3)

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

func findTwoDays(amounts []string, days int, bikeCost int) (res1 int, res2 int) {
	res2 = -1
	res1 = -1

	if days == 1 {
		a, _ := strconv.Atoi(amounts[0])
		if a >= bikeCost {
			res1 = 0
		}
		if a >= 2*bikeCost {
			res2 = 0
		}

		return
	}

	last, _ := strconv.Atoi(amounts[days-1])
	if last < bikeCost {
		return
	}

	res1 = findDay(amounts, 0, days-1, bikeCost)

	if last >= 2*bikeCost {
		res2 = findDay(amounts, res1, days-1, 2*bikeCost)
	}

	return
}

func findDay(amounts []string, start int, end int, bikeCost int) int {
	aEnd, _ := strconv.Atoi(amounts[end])
	if aEnd < bikeCost {
		return -1
	}

	if end == start {
		return start
	}

	if end-start == 1 {
		aStart, _ := strconv.Atoi(amounts[start])
		if aStart >= bikeCost {
			return start
		}

		return end
	}

	if end-start == 2 {
		aStart, _ := strconv.Atoi(amounts[start])
		if aStart >= bikeCost {
			return start
		}

		aStartOne, _ := strconv.Atoi(amounts[start+1])
		if aStartOne >= bikeCost {
			return start + 1
		}

		return end
	}

	aEndHalf, _ := strconv.Atoi(amounts[end/2])
	if aEndHalf >= bikeCost {
		return findDay(amounts, start, end/2, bikeCost)
	} else {
		return findDay(amounts, end/2+1, end, bikeCost)
	}
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
