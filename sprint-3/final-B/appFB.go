/*

посылка https://contest.yandex.ru/contest/23815/run-report/46455563/

-- ПРИНЦИП РАБОТЫ --
	Данное решение воспроизводит алгоритм in-place quick sort, уже описанный в задании.
	Для выбора опорного элемента мы используем метод medianOfThree - это позволяет существенно
уменьшить вероятность наткнуться на killer последовательность, в частности если исходный массив
изначально отсортирован. При таком подходе нужно дополнительно следить за тем, что если произошло перемещение
опорного элемента, то нужно соответсвенно менять и индекс-указатель на него.
Эту проверку мы делаем после вызова data.Swap(i, j).
	Для хранения опрного элемента мы не используем отдельной переменной - а храним лишь индекс-указатель
на элемент в массиве. Так сделано для того, чтобы использовать абстрактный метод Less(i,j), не зависящий
от конкретной имплементации интерфейса Sortable.
Эти идеи были подсмотрены в исходниках пакета sort

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
	На каждой стадии рекурсии мы делим исходный массив на две части, а затем меняем местами элементы таким образом,
что в конце получаем что в одной части все элементы меньше чем любой элемент другой части. Если это повторять
рекурсивно то в конце концов мы придем к тому что исходный массив полностью отсортируется.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
	Алгоритм quick sort довольно популярен и его сложность изучена и известна.
В худшем случае O(n*n) - это если неповезет с выбором опорного элемента настолько, что все partition стадии
будут выделять массыв лишь с одним элементом и по итогу мы на каждой итерции будем вынуждены проходить по всему массиву.
Но на практике все жатое случается крайне редко и в среднем все же мы на каждой итерации имеем дело с в 2 раза меньшим
числом элементов - т.е. в среднем O(n log n)

*/

package main

import (
	"bufio"
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

type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Entry struct {
	name string
	prob int
	fine int
}

type Leaderboard struct {
	data []Entry
}

type SortableInt []int

type Case struct {
	input    SortableInt
	expected SortableInt
}

func (c SortableInt) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c SortableInt) Swap(i int, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c SortableInt) Len() int {
	return len(c)
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	lb := readData(reader)
	lb.Sort()
	printLeaderBoard(lb, writer)
}

func quickSort(data Sortable, lo int, hi int) {
	if hi <= lo {
		return
	}

	p := partition(data, lo, hi)

	quickSort(data, lo, p)
	quickSort(data, p+1, hi)
}

func partition(data Sortable, lo int, hi int) int {
	p := (lo + hi) / 2
	medianOfThree(data, p, lo, hi)

	i := lo
	j := hi

	for {
		for ; data.Less(i, p); i++ {
		}

		for ; data.Less(p, j); j-- {
		}

		if i >= j {
			return j
		}

		data.Swap(i, j)

		oldM := p

		if i == oldM {
			p = j
		}

		if j == oldM {
			p = i
		}
	}
}

func medianOfThree(data Sortable, m1, m0, m2 int) {
	if data.Less(m1, m0) {
		data.Swap(m1, m0)
	}
	if data.Less(m2, m1) {
		data.Swap(m2, m1)
		if data.Less(m1, m0) {
			data.Swap(m1, m0)
		}
	}
}

func (lb *Leaderboard) Sort() {
	quickSort(lb, 0, lb.Len()-1)
}

func (lb *Leaderboard) Less(i, j int) bool {
	a := lb.data[i]
	b := lb.data[j]

	if a.prob != b.prob {
		return a.prob > b.prob
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

func printLeaderBoard(lb *Leaderboard, writer *bufio.Writer) {
	for _, entry := range lb.data {
		writer.WriteString(entry.name)
		writer.WriteString("\n")
	}

	writer.Flush()
}

func readData(reader *bufio.Reader) (lb *Leaderboard) {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanLines)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	lb = &Leaderboard{}

	for i := 0; i < n; i++ {
		sc.Scan()
		fields := strings.Fields(sc.Text())

		name := fields[0]
		prob, _ := strconv.Atoi(fields[1])
		fine, _ := strconv.Atoi(fields[2])

		lb.data = append(lb.data, Entry{name, prob, fine})
	}

	return
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}
