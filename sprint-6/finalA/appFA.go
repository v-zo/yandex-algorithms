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
	"container/heap"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

const errorMessage = "Oops! I did it again"

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

	graph := NewGraph(n, edges)
	totalWeight, err := findMST(graph)

	if err != nil {
		writeString(writer, err.Error())
	} else {
		writeString(writer, strconv.Itoa(totalWeight))
	}
}

type Vertex []Edge

func findMST(graph Graph) (totalWeight int, err error) {
	notAdded := graph.adjMap
	pq := PriorityQueue{}

	addVertex := func(v int) {
		delete(notAdded, v)
		for _, edge := range graph.edges {
			if edge.from == v && notAdded[edge.to] != nil {
				heap.Push(&pq, edge)
			}
		}
	}

	addVertex(1)
	for len(notAdded) > 0 && len(pq) > 0 {
		e := heap.Pop(&pq).(*Item).value
		if _, ok := notAdded[e.to]; ok {
			totalWeight += e.weight
			addVertex(e.to)
		}
	}

	if len(notAdded) != 0 {
		//goland:noinspection GoErrorStringFormat
		err = errors.New(errorMessage)
	}

	return
}

type Edge struct {
	from   int
	to     int
	weight int
}

type Graph struct {
	adjMap AdjacencyMap
	edges  []Edge
}

func NewGraph(size int, edges []Edge) Graph {
	adjMap := getAdjacencyMap(size, edges)

	return Graph{adjMap, edges}
}

type AdjacencyMap map[int]Vertex

func getAdjacencyMap(size int, edges []Edge) AdjacencyMap {
	m := len(edges)
	adjMap := make(AdjacencyMap)
	for i := 1; i <= size; i++ {
		adjMap[i] = Vertex{}
	}

	for i := 0; i < m; i++ {
		edg := edges[i]
		adjMap[edg.from] = append(adjMap[edg.from], edg)
	}

	return adjMap
}

func readData(reader io.Reader) (n int, outputEdges []Edge) {
	sc := bufio.NewScanner(reader)
	sc.Scan()
	firstLineData := toIntArray(sc.Text(), 2)
	n = firstLineData[0]

	uniqueEdges := scanUniqueEdges(func() (string, bool) {
		if sc.Scan() {
			return sc.Text(), true
		} else {
			return "", false
		}
	})

	for _, edge := range uniqueEdges {
		outputEdges = append(outputEdges, edge)
		revertedEdge := Edge{edge.to, edge.from, edge.weight}
		outputEdges = append(outputEdges, revertedEdge)
	}

	return
}

func scanUniqueEdges(next func() (string, bool)) (uniqueEdges map[int]Edge) {
	uniqueEdges = make(map[int]Edge)

	for txt, hasNext := next(); hasNext; txt, hasNext = next() {
		ed := toIntArray(txt, 3)
		hash := CantorPairingFunction(ed[0], ed[1])
		newWeight := ed[2]

		if edge, ok := uniqueEdges[hash]; ok {
			if newWeight > edge.weight {
				uniqueEdges[hash] = Edge{edge.from, edge.to, newWeight}
			}
		} else {
			if ed[0] != ed[1] {
				uniqueEdges[hash] = Edge{ed[0], ed[1], newWeight}
			}
		}
	}

	return
}

func CantorPairingFunction(num1, num2 int) int {
	var k1, k2 int
	if num1 < num2 {
		k1 = num1
		k2 = num2
	} else {
		k1 = num2
		k2 = num1
	}

	num := (k1 + k2) * (k1 + k2 + 1)

	return num/2 + k2
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

func writeString(writer io.Writer, s string) {
	_, err := io.WriteString(writer, s)
	check(err)
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

type Item struct {
	value Edge // The value of the item; arbitrary.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].value.weight > pq[j].value.weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := &Item{x.(Edge), n}
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
