/*

посылка ******

-- ПРИНЦИП РАБОТЫ --

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := openFile("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	Solve(reader, writer)
}

type Interface interface {
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
}

type Leaderboard struct {
	data []int
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	arr := []int{0, 0, 1, 0, 0, 1}
	var lb Interface = Leaderboard{arr}
	//lb =
	quickSort(lb, 0, 5)

	fmt.Println(arr)
	//writer.WriteString("\n")

	//writer.Flush()
}

func (lb Leaderboard) Less(i, j int) bool {
	return lb.data[i] < lb.data[j]
}

func (lb Leaderboard) Swap(i int, j int) {
	lb.data[i], lb.data[j] = lb.data[j], lb.data[i]
}

func (lb Leaderboard) Len() int {
	return len(lb.data)
}

func quickSort(data Interface, lo int, hi int) {
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

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}
