/*
- все упорядочены 1 2 3

- k в "большом" отрезке (k=5) 4 5 6 1 2 3
								^

- k в "малом" отрезке (k=2) 4 5 6 1 2 3
									^

- mid в "большом" отрезке (mid=1) 4 5 6 1 2 3
									^

- mid в "малом" отрезке (mid=4) 4 5 6 1 2 3
									    ^
*/

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

type Side int

const (
	undef Side = iota
	left
	right
)

func binarySearch(elements IntList, st int, end int, k int) int {
	stEl := elements.getEl(st)
	endEl := elements.getEl(end)

	if endEl == k {
		return end
	}

	if stEl == k {
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

	switch getSide(stEl, midEl, endEl, k) {
	case left:
		return binarySearch(elements, st, mid-1, k)
	case right:
		return binarySearch(elements, mid+1, end, k)
	default:
		return -1
	}
}

func getSide(stEl int, midEl int, endEl int, k int) (side Side) {
	if midEl < endEl { // lesser (right) segment
		if k < midEl {
			side = left
		} else {
			if k < endEl {
				side = right
			} else {
				side = left
			}
		}
	} else { // greater (left) segment
		if k > midEl {
			side = right
		} else {
			if k > stEl {
				side = left
			} else {
				side = right
			}
		}
	}

	if side == undef {
		panic("side is undefined")
	}

	return
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
