/* Doesnt work */

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

type Segment struct {
	lo int
	hi int
}

type FB struct {
	lows  []int
	highs []int
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	readPair := func() (segment Segment) {
		line := scanner.Text()
		fields := strings.Fields(line)
		start, _ := strconv.Atoi(fields[0])
		end, _ := strconv.Atoi(fields[1])
		segment = Segment{start, end}

		return
	}

	var fbs FB

	for scanner.Scan() {
		segment := readPair()
		fbs.insert(segment)
	}

	//for _, fb := range fbs.s {
	//	writer.WriteString(strconv.Itoa(fb.lo) + " " + strconv.Itoa(fb.hi))
	//	writer.WriteString("\n")
	//}

	writer.Flush()
}

func (fbs *FB) insert(seg Segment) {
	start := insertionIndex(fbs.highs, seg.hi)
	end := insertionIndex(fbs.lows, seg.lo)

	fbs.lows = add(fbs.lows, start, end, seg.lo)
	fbs.highs = add(fbs.highs, start, end, seg.hi)
}

func add(arr []int, start int, end int, el int) []int {
	if len(arr) == 0 {
		return []int{el}
	}

	arr = append(arr[:start], el)
	if start < len(arr)-1 {
		arr = append(arr, arr[end:]...)
	}

	return arr
}

func insertionIndex(arr []int, el int) int {
	if len(arr) == 0 {
		return 0
	}

	for i, item := range arr {
		if item == el {
			return i - 1
		}
		if item > el {
			return i
		}
	}

	return len(arr)
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
