/*

посылка https://contest.yandex.ru/contest/24414/problems/B/

-- ПРИНЦИП РАБОТЫ --
	Число корзин и формула для номера корзины  подобраны согласно формуле в конце главы "Выбор размера хеш-таблицы и
вычисление номера корзины"
https://praktikum.yandex.ru/learn/algorithms/courses/7f101a83-9539-4599-b6e8-8645c3f31fad/sprints/3003/topics/618173c7-3c0e-4955-b88b-d7146f9ffe2e/lessons/db4e40bc-75c2-4302-95ae-e9cc04f86546/
	Разрешать коллизии следует с помощью метода цепочек: в корзинах мы храним не сами значения,
а ссылку на связный список. Если при добавлении элемента возникает коллизия мы добавляем этот элемент
в начало списка.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
	Мы выбрали такую формулу для вычисления номера корзины, для которой известно что она обеспечивает хорошую
равномерность заполнения корзин.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
	Средняя сложность операций в хеш-таблице равняется O(1+α), где α -  отношение общего числа запросов к
числу корзин, в нашем случае 1 000 000/65536 = 15.2. Это выше чем рекомендуемое значение
1 ≤ α ≤ 2, цепочки получаются длинными, но с поставленной задачи алгоритм справляется.
*/

package main

import (
	"bufio"
	"container/list"
	"os"
	"strconv"
	"strings"
)

const (
	s        = 2654435769
	twoPow32 = 4294967296
)

func main() {
	file := openFile("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	Solve(reader, writer)
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	yaReader := &YaReader{reader}
	n := yaReader.readInt()
	sc := bufio.NewScanner(yaReader)
	sc.Split(bufio.ScanLines)

	ht := createHashTable(16) // 2^16==65536 < 100000 < 2^17==131072
	ex := &Executor{ht, writer}

	for i := 0; i < n; i++ {
		sc.Scan()
		ex.run(sc.Text())
	}

	writer.Flush()
}

type Entry struct {
	key int
	val string
}

type HashTable struct {
	data []*list.List
	m    int
}

func createHashTable(p int) *HashTable {
	arr := make([]*list.List, powInt(2, p))
	for i := range arr {
		arr[i] = list.New()
	}
	return &HashTable{arr, 32 - p}
}

func (q *HashTable) bucket(k int) *list.List {
	return q.data[(k*s%twoPow32)>>q.m]
}

func (q *HashTable) put(key int, val string) {
	l := q.bucket(key)
	e := findByKey(l, key)

	if e != nil {
		l.Remove(e)
	}

	l.PushFront(&Entry{key, val})
}

func (q *HashTable) get(key int) string {
	l := q.bucket(key)
	e := findByKey(l, key)

	if e == nil {
		return "None"
	}

	return e.Value.(*Entry).val
}

func (q *HashTable) delete(key int) string {
	l := q.bucket(key)
	e := findByKey(l, key)

	if e == nil {
		return "None"
	}

	val := e.Value.(*Entry).val
	l.Remove(e)

	return val
}

func findByKey(l *list.List, key int) *list.Element {
	e := l.Front()
	for e != nil && e.Value.(*Entry).key != key {
		e = e.Next()
	}

	return e
}

type Executor struct {
	hashTable *HashTable
	writer    *bufio.Writer
}

func (ex *Executor) writeLn(s string) {
	ex.writer.WriteString(s)
	ex.writer.WriteString("\n")
}

func (ex *Executor) run(command string) {
	fields := strings.Fields(command)

	if fields[0] == "put" {
		key := toInt(fields[1])
		val := fields[2]
		ex.hashTable.put(key, val)
	}

	if fields[0] == "get" {
		key := toInt(fields[1])
		val := ex.hashTable.get(key)
		ex.writeLn(val)
	}

	if fields[0] == "delete" {
		key := toInt(fields[1])
		val := ex.hashTable.delete(key)
		ex.writeLn(val)
	}
}

func toInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
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

type YaReader struct {
	*bufio.Reader
}

func (reader *YaReader) readString() string {
	line, _ := reader.ReadString('\n')
	return strings.TrimRight(line, "\n")
}

func (reader *YaReader) readInt() int {
	line, _ := reader.ReadString('\n')
	res, _ := strconv.Atoi(strings.TrimRight(line, "\n"))
	return res
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}
