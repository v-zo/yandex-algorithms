/*
~Doesnt work
*/

package main

import (
	"bufio"
	"math"
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
	s []Segment
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
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
		fbs.intersect(segment, 0)
	}

	for _, fb := range fbs.s {
		writer.WriteString(strconv.Itoa(fb.lo) + " " + strconv.Itoa(fb.hi))
		writer.WriteString("\n")
	}

	writer.Flush()
}

func (fbs *FB) intersect(seg Segment, st int) *FB {
	if len(fbs.s) == st {
		fbs.s = append(fbs.s, seg)
		return fbs
	}

	for i := st; i < len(fbs.s); i++ {
		rel := compare(seg, fbs.s[i])

		if rel == equal {
			return fbs
		}

		if rel == more {
			newFbs := append(fbs.s[:i+1], seg)
			fbs.s = append(newFbs, fbs.s[i+1:]...)
			fbs.s = fbs.intersect(seg, i+1).s

			return fbs
		}

		if rel == intersect {
			newSeg := join(fbs.s[i], seg)

			fbs.s[i] = newSeg
			fbs.s = fbs.intersect(newSeg, i).s

			return fbs
		}
	}

	fbs.s = append([]Segment{seg}, fbs.s...)
	return fbs
}

type Rel int

const (
	less Rel = iota
	more
	intersect
	equal
)

func join(a Segment, b Segment) Segment {
	return Segment{min(a.lo, b.lo), max(a.hi, b.hi)}
}

func compare(a Segment, b Segment) Rel {
	if a.hi == b.hi && a.lo == b.lo {
		return equal
	}

	if a.hi < b.lo {
		return less
	}

	if a.lo > b.hi {
		return more
	}

	return intersect
}

func max(nums ...int) (result int) {
	result = math.MinInt64
	for _, num := range nums {
		if num > result {
			result = num
		}
	}

	return
}

func min(nums ...int) (result int) {
	result = math.MaxInt64
	for _, num := range nums {
		if num < result {
			result = num
		}
	}

	return
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
