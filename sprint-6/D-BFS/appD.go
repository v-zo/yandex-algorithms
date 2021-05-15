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

	SplitToWriter(writer, MainDFS(n, edges, start), " ")
}

func MainDFS(size int, edges [][]int, start int) []int {
	adj := NewGraph(size, edges)
	res := &[]int{start}
	adj.BFS(start, res)

	return *res
}

func (g *Graph) BFS(v int, res *[]int) {
	g.colors[v] = gray
	neighbours := g.adjMap[v]

	sort.Slice(neighbours, func(i, j int) bool {
		return neighbours[i] < neighbours[j]
	})

	for _, w := range neighbours {
		if g.colors[w] == white {
			*res = append(*res, w)
			g.BFS(w, res)
		}
	}

	g.colors[v] = black

	return
}

type Graph struct {
	adjMap AdjacencyMap
	len    int
	colors map[int]Color
}

func NewGraph(size int, edges [][]int) Graph {
	adjMap, colors := getAdjMap(edges)

	return Graph{adjMap, size, colors}
}

type Queue struct {
	ch chan int
}

func (q *Queue) Init() {
	q.ch = make(chan int, 300)
}

func (q *Queue) Push(value int) {
	q.ch <- value
}

func (q *Queue) Pop() int {
	return <-q.ch
}

type AdjacencyMap map[int][]int

func getAdjMap(edges [][]int) (AdjacencyMap, map[int]Color) {
	m := len(edges)
	adjMap := make(AdjacencyMap)
	colors := make(map[int]Color, m)

	for i := 0; i < m; i++ {
		dot1, dot2 := edges[i][0], edges[i][1]

		adjMap[dot1] = append(adjMap[dot1], dot2)
		if _, ok := colors[dot1]; !ok {
			colors[dot1] = white
		}

		adjMap[dot2] = append(adjMap[dot2], dot1)
		if _, ok := colors[dot2]; !ok {
			colors[dot2] = white
		}
	}

	return adjMap, colors
}

func SplitToWriter(writer io.Writer, a []int, sep string) {
	if len(a) == 0 {
		return
	}

	writeNumber(writer, a[0])
	for i := 1; i < len(a); i++ {
		writeString(writer, sep)
		writeNumber(writer, a[i])
	}

	writeString(writer, "\n")
}

func writeNumber(writer io.Writer, num int) {
	s := strconv.Itoa(num)
	writeString(writer, s)
}

func writeString(writer io.Writer, s string) {
	_, err := io.WriteString(writer, s)
	check(err)
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
