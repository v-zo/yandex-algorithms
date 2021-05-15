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
	graph := NewGraph(n, edges)
	mst, err := findMST(graph, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mst)
	}
}

type Vertex []Edge

func findMST(graph Graph, start int) (mst []Edge, err error) {
	var added []int
	notAdded := graph.adjMap
	var mstEdges []Edge

	addVertex := func(v int) {
		added = append(added, v)
		delete(notAdded, v)
		for _, edge := range graph.edges {
			if edge.from == v && notAdded[edge.to] != nil {
				mstEdges = append(mstEdges, edge)
			}
		}
	}

	addVertex(start)
	for len(notAdded) > 0 && len(mstEdges) > 0 {
		e := extractMaximum(&mstEdges)
		if _, ok := notAdded[e.to]; ok {
			mst = append(mst, e)
			addVertex(e.to)
		}
	}

	if len(notAdded) != 0 {
		//goland:noinspection GoErrorStringFormat
		err = errors.New("Oops! I did it again")
	}

	return
}

func extractMaximum(edges *[]Edge) Edge {
	maxKey := 0
	max := (*edges)[maxKey]
	for key, val := range *edges {
		if val.weight > max.weight {
			max = val
			maxKey = key
		}
	}

	*edges = append((*edges)[:maxKey], (*edges)[maxKey+1:]...)

	return max

}

type Edge struct {
	from   int
	to     int
	weight int
}

type Graph struct {
	adjMap AdjacencyMap
	size   int
	edges  []Edge
}

func NewGraph(size int, edges []Edge) Graph {
	adjMap := getAdjMap(edges)

	return Graph{adjMap, size, edges}
}

type AdjacencyMap map[int]Vertex

func getAdjMap(edges []Edge) AdjacencyMap {
	m := len(edges)
	adjMap := make(AdjacencyMap)

	for i := 0; i < m; i++ {
		edg := edges[i]
		adjMap[edg.from] = append(adjMap[edg.from], edg)
		reversedEdg := Edge{edg.to, edg.from, edg.weight}
		adjMap[edg.to] = append(adjMap[edg.to], reversedEdg)
	}

	return adjMap
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

func createPriorityQueue(edges []Edge) PriorityQueue {
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(edges))
	for i, edge := range edges {
		pq[i] = &Item{
			value: edge,
			index: i,
		}
	}
	heap.Init(&pq)

	return pq
}

// An Item is something we manage in a priority queue.
type Item struct {
	value Edge // The value of the item; arbitrary.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
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
	item := x.(*Item)
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

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value Edge) {
	item.value = value
	heap.Fix(pq, item.index)
}
