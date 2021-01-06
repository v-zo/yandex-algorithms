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

type IntList struct {
	str []string
}

func (list IntList) getEl(i int) (value int) {
	value, _ = strconv.Atoi(list.str[i])

	return
}

func binarySearch(elements IntList, st int, end int, k int) int {
	if elements.getEl(end) == k {
		return end
	}
	if elements.getEl(st) == k {
		return st
	}

	if end-st <= 1 {
		return -1
	}

	mid := (st + end) / 2
	midEl := elements.getEl(mid)

	if midEl == k {
		return mid
	}

	side := "left"
	if midEl < elements.getEl(end) { // lesser
		if midEl > k {
			if k > elements.getEl(st) {

			} else {
				side = "right"
			}
		} else {
			side = "right"
		}
	} else { // greater
		if midEl > k {
			if k > elements.getEl(st) {

			} else {
				side = "right"
			}
		} else {
			side = "right"
		}
	}

	if side == "left" {
		return binarySearch(elements, st, mid-1, k)
	} else {
		return binarySearch(elements, mid+1, end, k)
	}

	//if midEl > k {
	//	if midEl < elements.getEl(end) {
	//		return binarySearch(elements, mid+1, end, k)
	//	}
	//	return binarySearch(elements, mid+1, end, k)
	//} else {
	//	if midEl > elements.getEl(end) {
	//		return binarySearch(elements, st, mid-1, k)
	//	}
	//	return binarySearch(elements, mid+1, end, k)
	//}
}

func readData(reader *bufio.Reader) (n int, k int, elements IntList) {
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line3, _ := reader.ReadString('\n')

	n, _ = strconv.Atoi(strings.TrimRight(line1, "\n"))
	k, _ = strconv.Atoi(strings.TrimRight(line2, "\n"))
	elements = IntList{strings.Fields(line3)}

	return
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
