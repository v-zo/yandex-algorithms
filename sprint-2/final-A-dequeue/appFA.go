/*
посылка 45866566

-- ПРИНЦИП РАБОТЫ --
Данный дек работает по тому же самому принципу что и очередь на массиве,
реализация которой хорошо описана в теоретической части курса (
	https://praktikum.yandex.ru/learn/algorithms/courses/7f101a83-9539-4599-b6e8-8645c3f31fad/sprints/4526/topics/3fe165ac-9a25-44db-b5bf-106709d1c140/lessons/b5036361-6d27-45be-ac0a-3946d9a0fcde/
).
Добавлена лишь возможность перемещать указатель хвоста, а не только головы.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Если допустить что алгоритм очереди на массиве из теоретической части курса корректен,
то корректна и данная реализация дека, т.к. используется тот же самый алгоритм.
Операции с хвостом в нашем деке - по сути копии опреций с головой в очереди на массиве.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
каждая из 4 операций дека сводятся к перемещению указателя.
Перемещение указателя - арифметическая операция, следовательно сложность любой из 4 операций равна O(1)
*/

package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Dequeue struct {
	data []string
	cap  int
	head int
	tail int
	sz   int
}

func main() {
	file := openFile("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	Solve(reader, writer)
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	queue := createDequeue(n)
	ex := &Executor{queue, writer}

	for scanner.Scan() {
		ex.run(scanner.Text())
	}

	writer.Flush()
}

type Executor struct {
	dequeue *Dequeue
	writer  *bufio.Writer
}

func (ex *Executor) writeLn(s string) {
	ex.writer.WriteString(s)
	ex.writer.WriteString("\n")
}

func (ex *Executor) run(command string) {
	if strings.Contains(command, "push_back") {
		value := strings.Split(command, " ")[1]
		err := ex.dequeue.pushBack(value)

		if err != nil {
			ex.writeLn(err.Error())
		}
	}

	if strings.Contains(command, "push_front") {
		value := strings.Split(command, " ")[1]
		err := ex.dequeue.pushFront(value)

		if err != nil {
			ex.writeLn(err.Error())
		}
	}

	if command == "pop_front" {
		el, err := ex.dequeue.popFront()

		if err == nil {
			ex.writeLn(el)
		} else {
			ex.writeLn(err.Error())
		}
	}

	if command == "pop_back" {
		el, err := ex.dequeue.popBack()

		if err == nil {
			ex.writeLn(el)
		} else {
			ex.writeLn(err.Error())
		}
	}
}

func createDequeue(n int) *Dequeue {
	arr := make([]string, n, n)
	return &Dequeue{data: arr, cap: n, head: 1, tail: 0, sz: 0}
}

func (q *Dequeue) pushFront(x string) error {
	if q.sz == q.cap {
		return errors.New("error")
	}

	q.data[q.head] = x
	q.head = increment(q.head, q.cap)
	q.sz = q.sz + 1

	return nil
}

func (q *Dequeue) pushBack(x string) error {
	if q.sz == q.cap {
		return errors.New("error")
	}

	q.data[q.tail] = x
	q.tail = decrement(q.tail, q.cap)
	q.sz = q.sz + 1

	return nil
}

func (q *Dequeue) popBack() (el string, err error) {
	if q.sz == 0 {
		err = errors.New("error")

		return
	}

	q.tail = increment(q.tail, q.cap)
	el = q.data[q.tail]
	q.sz = q.sz - 1

	return
}

func (q *Dequeue) popFront() (el string, err error) {
	if q.sz == 0 {
		err = errors.New("error")

		return
	}

	q.head = decrement(q.head, q.cap)
	el = q.data[q.head]
	q.sz = q.sz - 1

	return
}

func decrement(x int, cap int) int {
	return (cap + (x-1)%cap) % cap
}

func increment(x int, cap int) int {
	return (x + 1) % cap
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
