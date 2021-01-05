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
	size := (n+m)/2 + 1

	scanNum := func(scanner *bufio.Scanner) (num int) {
		scanner.Scan()
		num, _ = strconv.Atoi(scanner.Text())

		return
	}

	left := scanNum(north)
	right := scanNum(south)

	var merged []int

	i := 0
	j := 0
	for i+j+1 <= size {
		if left < right {
			merged = append(merged, left)
			left = scanNum(north)
			i++
		} else {
			merged = append(merged, right)
			right = scanNum(south)
			j++
		}

		if i == n {
			tail := size - i - j
			for k := 0; k < tail; k++ {
				merged = append(merged, right)
				right = scanNum(south)
			}

			break
		}

		if j == m {
			tail := size - i - j
			for k := 0; k < tail; k++ {
				merged = append(merged, left)
				left = scanNum(north)
			}

			break
		}
	}

	if (n+m)%2 == 0 {
		output = getEvenOutput(merged[size-1] + merged[size-2])
	} else {
		output = getOddOutput(merged[size-1])
	}

	return
}

func getEvenOutput(num int) string {
	s2 := strconv.Itoa(num / 2)
	if num%2 == 1 {
		return s2 + "." + "5"
	} else {
		return s2
	}
}

func getOddOutput(num int) string {
	return strconv.Itoa(num)
}

func readData(reader *bufio.Reader) (n int, m int, north *bufio.Scanner, south *bufio.Scanner) {
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line3, _ := reader.ReadString('\n')
	line4, _ := reader.ReadString('\n')

	n, _ = strconv.Atoi(strings.TrimRight(line1, "\n"))
	m, _ = strconv.Atoi(strings.TrimRight(line2, "\n"))

	createScanner := func(line string) *bufio.Scanner {
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
