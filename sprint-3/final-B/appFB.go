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
	"sort"
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

type Entry struct {
	name string
	prob int
	fine int
}

type Leaderboard struct {
	data []Entry
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	lb := readData(reader)
	lb.Sort()
	printLeaderBoard(lb, writer)
}

func quickSort(data sort.Interface, lo int, hi int) {
	if hi-lo <= 1 {
		if data.Less(hi, lo) {
			data.Swap(lo, hi)
		}

		return
	}

	m := (lo + hi) / 2

	i := lo
	j := hi

	for {
		for ; data.Less(i, m) && i < j-1; i++ {
		}
		for ; !data.Less(j, m) && i < j-1; j-- {
		}

		if j-i == 1 {
			break
		}

		data.Swap(i, j)
	}

	quickSort(data, lo, m)
	quickSort(data, m, hi)
}

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
