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

func NewColors(m int) []Color {
	colors := make([]Color, m)
	for i := 0; i < m; i++ {
		colors[i] = white
	}

	return colors
}

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

	return adj.DFS(start)
}

func (adj *AdjList) DFS(v int) (res []int) {
	adj.colors[v-1] = gray
	res = append(res, v)

	for _, w := range adj.adjMap[v] {
		if adj.colors[w-1] == white {
			res = append(res, adj.DFS(w)...)
		}
	}

	adj.colors[v-1] = black
	delete(adj.adjMap, v)

	return
}

type AdjList struct {
	adjMap AdjMap
	len    int
	colors []Color
}

func NewAdjList(size int, edges [][]int) AdjList {
	adjMap := getAdjMap(edges)
	colors := NewColors(size)

	return AdjList{adjMap, size, colors}
}

type AdjMap map[int][]int

func getAdjMap(edges [][]int) AdjMap {
	m := len(edges)
	adjMap := make(AdjMap)

	for i := 0; i < m; i++ {
		dot1, dot2 := edges[i][0], edges[i][1]

		adjMap[dot1] = append(adjMap[dot1], dot2)
		adjMap[dot2] = append(adjMap[dot2], dot1)
	}

	for _, v := range adjMap {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}

	return adjMap
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
