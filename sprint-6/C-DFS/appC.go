package main

import (
	"bufio"
	"io"
	"os"
	"sort"
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

type Color int

const (
	white Color = iota
	gray
	black
)

func Solve(reader *bufio.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	yaScanner := &YaScanner{scanner}

	n, m := yaScanner.scanPair()

	var edges [][]int
	for i := 0; i < m; i++ {
		dot1, dot2 := yaScanner.scanPair()
		edges = append(edges, []int{dot1, dot2})
	}

	yaScanner.Scan()
	start, _ := strconv.Atoi(yaScanner.Text())

	outLine := SplitToString(MainDFS(n, edges, start), " ")
	io.WriteString(writer, outLine)
}

func MainDFS(size int, edges [][]int, start int) []int {
	adj := NewAdjList(size, edges)
	res := &[]int{start}
	adj.DFS(start, res)

	return *res
}

func (adj *AdjList) DFS(v int, res *[]int) {
	adj.colors[v] = gray
	neighbours := adj.adjMap[v]

	sort.Slice(neighbours, func(i, j int) bool {
		return neighbours[i] < neighbours[j]
	})

	for _, w := range neighbours {
		if adj.colors[w] == white {
			*res = append(*res, w)
			adj.DFS(w, res)
		}
	}

	adj.colors[v] = black

	return
}

type AdjList struct {
	adjMap AdjMap
	len    int
	colors map[int]Color
}

func NewAdjList(size int, edges [][]int) AdjList {
	adjMap, colors := getAdjMap(edges)

	return AdjList{adjMap, size, colors}
}

type AdjMap map[int][]int

func getAdjMap(edges [][]int) (AdjMap, map[int]Color) {
	m := len(edges)
	adjMap := make(AdjMap)
	colors := make(map[int]Color, m)

	for i := 0; i < m; i++ {
		dot1, dot2 := edges[i][0], edges[i][1]

		adjMap[dot1] = append(adjMap[dot1], dot2)
		adjMap[dot2] = append(adjMap[dot2], dot1)

		if _, ok := colors[dot1]; !ok {
			colors[dot1] = white
		}
		if _, ok := colors[dot2]; !ok {
			colors[dot2] = white
		}
	}

	return adjMap, colors
}

func SplitToString(a []int, sep string) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}

	res := strings.Join(b, sep)
	return res + "\n"
}

type YaScanner struct {
	*bufio.Scanner
}

func (scanner *YaScanner) scanPair() (int, int) {
	scanner.Scan()
	line := scanner.Text()
	fields := strings.Fields(line)
	n, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])

	return n, m
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
