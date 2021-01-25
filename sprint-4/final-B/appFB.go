/*

посылка
https://contest.yandex.ru/contest/24414/run-report/47302058/

-- ПРИНЦИП РАБОТЫ --
	Данное решение реализует стандартную структуру данных "Хэш Таблица" с целочисленными ключами и
строковыми значениями.
	Число корзин и формула для номера корзины подобраны согласно формуле в конце главы "Выбор размера хеш-таблицы и
вычисление номера корзины":
bucket(h)=(h⋅s mod 2^32)≫(32−p), где s=2654435769. Здесь p - целое число, логарифм по основанию 2 от числа корзин.
При выборе значения параметра p мы должны учитывать что по условию задачи число корзин не должно превышать 100000.
2^16==65536 < 100000 < 2^17==131072, т.е. 65536 - степень 2, максимально допустимая ограничением. Поэтому в качестве p
мы выбрали значение 16.
	Коллизии разрешаем с помощью метода цепочек: в корзинах мы храним не сами значения, а ссылку на связный список.
Если при добавлении элемента возникает коллизия мы добавляем этот элементт в начало списка.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
	Реализованная нами структура данных поддерживает операции чтения, записи и удаления элементов по ключу - что
и требовалось исполнить по условию задачи.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
	Средняя сложность операций в хеш-таблице равняется O(1+α), где α -  отношение общего числа запросов к
числу корзин, в нашем случае 1 000 000/65536 = 15.2. Это выше чем рекомендуемое значение
1 ≤ α ≤ 2, цепочки получаются длинными, но с поставленной задачи алгоритм справляется.

*/

package main

import (
	"bufio"
	"container/list"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	goldenMultiplier = 2654435769
	twoPow32         = 4294967296
	log2BucketSize   = 16
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

func Solve(reader *bufio.Reader, writer io.Writer) {
	yaReader := &YaReader{reader}
	n := yaReader.readInt()
	sc := bufio.NewScanner(yaReader)

	ht := createHashTable(log2BucketSize)
	ex := &Executor{ht, writer}

	for i := 0; i < n; i++ {
		sc.Scan()
		ex.run(sc.Text())
	}
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
	size := int(math.Pow(2, float64(p)))
	arr := make([]*list.List, size)
	for i := range arr {
		arr[i] = list.New()
	}
	return &HashTable{arr, 32 - p}
}

func (q *HashTable) bucket(k int) *list.List {
	return q.data[(k*goldenMultiplier%twoPow32)>>q.m]
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
	writer    io.Writer
}

func (ex *Executor) writeLn(s string) {
	_, err := io.WriteString(ex.writer, s)
	check(err)
	_, err = io.WriteString(ex.writer, "\n")
	check(err)
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

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

type YaReader struct {
	*bufio.Reader
}

func (reader *YaReader) readString() string {
	line, err := reader.ReadString('\n')
	check(err)
	return strings.TrimRight(line, "\n")
}

func (reader *YaReader) readInt() int {
	line, err := reader.ReadString('\n')
	check(err)
	res, err := strconv.Atoi(strings.TrimRight(line, "\n"))
	check(err)
	return res
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
