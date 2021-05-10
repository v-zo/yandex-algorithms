/*

посылка 51374855

-- ПРИНЦИП РАБОТЫ --
	Алгоритм состоит условно из двух этапов:
1 - строим невозрастающую кучу из исходного массива.
2 - начиная с конца меняем местами элементы с вершиной, попутно восстанавливая свойства кучи просеиванием вниз
	На втором этапе при обходе массива с конца мы меняем местами текущий элемент с головным элементом кучи. Назовем
условно "хвостом" ту часть массива, куда перемещаются элементы из вершины кучи.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
	Т.к. вершина кучи - максималый элемент, то при перестановке мы присоединяем к началу "хвоста" массива наибольший для
оставшейся (неотсортированной) части элемент, но меньший чем текущие элементы "хвоста".
Таким образом, элементы "хвоста" отсортированы.
В результате итераций по всем элементам массива наступит момент, когда "хвостом" будет являться весь массив целиком,
а значит он будет отсортирован, что и требовалось доказать.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
	Каждый из этапов совершает одинаковые по сложности процедуры: порядка n раз вызывается функция просеивания.
Сам алгоритм просеивания имеет сложность O(log n) (верхняя оценка), следовательно итоговая сложность: O(n * log n)

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
	Данный алгоритм мутирует исходный массив "in place". Эта идея была подсмотрена в стандартном пакете sort.
Т.к. дополнительный массив создавать не требуется, сложность по дополнительной памяти состаляет O(1)

*/

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

func Solve(reader io.Reader, writer io.Writer) {
	lb := readData(reader)
	lb.Sort()
	lb.Print(writer)
}

func heapSort(data sort.Interface) {
	buildHeap(data)

	for i := data.Len() - 1; i >= 0; i-- {
		data.Swap(0, i)
		siftDown(data, 0, i)
	}
}

func buildHeap(data sort.Interface) {
	size := data.Len()

	for i := (size - 1) / 2; i >= 0; i-- {
		siftDown(data, i, size)
	}
}

func siftDown(data sort.Interface, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(child, child+1) {
			child++
		}
		if !data.Less(root, child) {
			return
		}
		data.Swap(root, child)
		root = child
	}
}

type Entry struct {
	name   string
	points int
	fine   int
}

type Leaderboard struct {
	data []Entry
	sort.Interface
}

func (lb *Leaderboard) Sort() {
	heapSort(lb)
}

func (lb *Leaderboard) Less(i, j int) bool {
	a := lb.data[i]
	b := lb.data[j]

	if a.points != b.points {
		return a.points > b.points
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

func (lb *Leaderboard) Print(writer io.Writer) {
	var err error

	for _, entry := range lb.data {
		_, err = io.WriteString(writer, entry.name+"\n")
		check(err)
	}
}

func readData(reader io.Reader) (lb *Leaderboard) {
	sc := bufio.NewScanner(reader)
	sc.Scan()
	n := atoi(sc.Text())

	lb = &Leaderboard{}

	for i := 0; i < n; i++ {
		sc.Scan()
		entry := parseEntry(sc.Text())
		lb.data = append(lb.data, entry)
	}

	return
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	check(err)

	return n
}

func parseEntry(s string) Entry {
	fields := strings.Fields(s)
	name := fields[0]
	prob := atoi(fields[1])
	fine := atoi(fields[2])

	return Entry{name, prob, fine}
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
