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

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	arr := []int{0, 0, 1, 0, 0, 1}
	quickSort(arr, 0, 5)

	fmt.Println(arr)
	//writer.WriteString("\n")

	//writer.Flush()
}

func quickSort(arr []int, lo int, hi int) {
	if hi-lo <= 1 {
		if arr[lo] > arr[hi] {
			swap(arr, lo, hi)
		}

		return
	}

	m := (lo + hi) / 2
	pivot := arr[m]

	i := lo
	j := hi

	for {
		for ; arr[i] < pivot && i < j-1; i++ {
		}
		for ; !(arr[j] < pivot) && i < j-1; j-- {
		}

		if j-i == 1 {
			break
		}

		swap(arr, i, j)
	}

	quickSort(arr, lo, m)
	quickSort(arr, m, hi)
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
