/*

посылка --

-- ПРИНЦИП РАБОТЫ --

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := openFile("input.txt")
	defer file.close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	Solve(reader, writer)

	err := writer.Flush()
	check(err)
}

func Solve(reader io.Reader, writer io.Writer) {
	n, edges := readData(reader)
	fmt.Println(n)
	fmt.Print(edges)
}

type Edge struct {
	from   int
	to     int
	weight int
}

func readData(reader io.Reader) (n int, edges []Edge) {
	sc := bufio.NewScanner(reader)
	sc.Scan()
	firstLineData := toIntArray(sc.Text(), 2)
	n = firstLineData[0]
	m := firstLineData[1]

	for i := 0; i < m; i++ {
		sc.Scan()
		ed := toIntArray(sc.Text(), 3)
		edges = append(edges, Edge{ed[0], ed[1], ed[2]})
	}

	return
}

func toIntArray(s string, size int) (res []int) {
	lineData := strings.Split(s, " ")

	for i := 0; i < size; i++ {
		res = append(res, atoi(lineData[i]))
	}

	return
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	check(err)

	return n
}

type File struct {
	*os.File
}

func openFile(path string) *File {
	osFile, err := os.Open(path)
	check(err)

	return &File{osFile}
}

func (file *File) close() {
	err := file.Close()
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
