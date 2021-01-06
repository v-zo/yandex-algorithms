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
	n, k, elements := readData(reader)

	result := binarySearch(elements, 0, n-1, k)

	writer.WriteString(strconv.Itoa(result))
	writer.WriteString("\n")

	writer.Flush()
}

func binarySearch(elements []string, st int, end int, k string) int {
	if end == st+1 {
		if elements[end] == k {
			return end
		}
		if elements[st] == k {
			return st
		}

		return -1
	}

	mid := (st + end) / 2

	if elements[mid] == k {
		return mid
	}

	if elements[mid] > elements[end] {
		if elements[mid] > k {
			return binarySearch(elements, st, mid-1, k)
		} else {
			return binarySearch(elements, mid+1, end, k)
		}
	} else {
		if elements[mid] < k {
			return binarySearch(elements, mid+1, end, k)
		} else {
			return binarySearch(elements, st, mid-1, k)
		}
	}
}

func readData(reader *bufio.Reader) (n int, k string, elements []string) {
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line3, _ := reader.ReadString('\n')

	n, _ = strconv.Atoi(strings.TrimRight(line1, "\n"))
	k = strings.TrimRight(line2, "\n")
	elements = strings.Fields(line3)

	return
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
