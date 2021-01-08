package main

import (
	"bufio"
	"fmt"
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
	//n, k := readData(reader)
	//result := fibonacciModulo(n, k, 1, 0)
	//
	//writer.WriteString(strconv.Itoa(result))
	//writer.WriteString("\n")
	//
	//writer.Flush()

	//A := [][]int{[]int{1,2},[]int{3,4}}
	//B := [][]int{[]int{2,0},[]int{1,2}}
	//fmt.Println(product(A,B))

	fmt.Println(powInt(2, 8))
}

func powInt(x, y int) (p int) {
	p = 1

	for y != 0 {
		if y%2 != 0 {
			p *= x
		}

		x *= x
		y /= 2
	}

	return
}

func fibonacciModulo(n, k, p, p0 int) int {
	for i := 0; i < n; i++ {
		p, p0 = p0, p
		p = (p + p0) % powInt(10, k)
	}

	return p
}

func product(A, B [][]int) (prod [][]int) {
	n := len(A)

	prod = make([][]int, n)
	for i := range prod {
		prod[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				prod[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return
}

func readData(reader *bufio.Reader) (n, k int) {
	line1, _ := reader.ReadString('\n')
	fields := strings.Fields(strings.TrimRight(line1, "\n"))
	n, _ = strconv.Atoi(fields[0])
	k, _ = strconv.Atoi(fields[1])

	return
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}
