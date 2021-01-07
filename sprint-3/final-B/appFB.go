/*

посылка ******

-- ПРИНЦИП РАБОТЫ --

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

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

type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Entry struct {
	name string
	prob int
	fine int
}

type Leaderboard struct {
	data []Entry
}

///////
//type SortableInt []int
//
//type Case struct {
//	input    SortableInt
//	expected SortableInt
//}
//
//func (c SortableInt) Less(i, j int) bool {
//	return c[i] < c[j]
//}
//
//func (c SortableInt) Swap(i int, j int) {
//	c[i], c[j] = c[j], c[i]
//}
//
//func (c SortableInt) Len() int {
//	return len(c)
//}
//
/////////

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	lb := readData(reader)
	lb.Sort()
	//si := SortableInt{3, 5, 10, 4, 1}
	//si := SortableInt{3, 2, 4, 1}
	//
	//quickSort(si, 0, 3)
	//fmt.Println(si)

	//printLeaderBoard(lb, writer)
}

func quickSort(data Sortable, lo int, hi int) {
	if hi <= lo {
		return
	}

	p := partition(data, lo, hi)

	quickSort(data, lo, p)
	quickSort(data, p+1, hi)
}

func partition(data Sortable, lo int, hi int) int {
	p := (lo + hi) / 2

	i := lo
	j := hi

	for {
		for ; data.Less(i, p); i++ {
		}

		for ; data.Less(p, j); j-- {
		}

		if i >= j {
			return j
		}

		data.Swap(i, j)

		oldM := p

		if i == oldM {
			p = j
		}

		if j == oldM {
			p = i
		}
	}
}

//func medianOfThree(data Sortable, m1, m0, m2 int) {
//	if data.Less(m1, m0) {
//		data.Swap(m1, m0)
//	}
//	if data.Less(m2, m1) {
//		data.Swap(m2, m1)
//		if data.Less(m1, m0) {
//			data.Swap(m1, m0)
//		}
//	}
//}

func (lb *Leaderboard) Sort() {
	quickSort(lb, 0, lb.Len()-1)
}

func (lb *Leaderboard) Less(i, j int) bool {
	a := lb.data[i]
	b := lb.data[j]

	if a.prob != b.prob {
		return a.prob > b.prob
	}

	if a.fine != b.fine {
		return a.fine < b.fine
	}

	return a.name < b.name
}

func (lb *Leaderboard) Swap(i int, j int) {
	lb.data[i], lb.data[j] = lb.data[j], lb.data[i]
}

func (lb *Leaderboard) Len() int {
	return len(lb.data)
}

func printLeaderBoard(lb *Leaderboard, writer *bufio.Writer) {
	for _, entry := range lb.data {
		writer.WriteString(entry.name)
		writer.WriteString("\n")
	}

	writer.Flush()
}

func readData(reader *bufio.Reader) (lb *Leaderboard) {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanLines)
	sc.Scan()

	lb = &Leaderboard{}

	for sc.Scan() {
		fields := strings.Fields(sc.Text())

		name := fields[0]
		prob, _ := strconv.Atoi(fields[1])
		fine, _ := strconv.Atoi(fields[2])

		lb.data = append(lb.data, Entry{name, prob, fine})
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
