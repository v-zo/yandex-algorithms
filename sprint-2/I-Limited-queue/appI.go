package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
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
	queue := createQueue(n)
	ex := &Executor{queue, writer}

	for scanner.Scan() {
		ex.run(scanner.Text())
	}

	writer.Flush()
}

type Executor struct {
	queue  *Queue
	writer *bufio.Writer
}

func (ex *Executor) writeLn(s string) {
	ex.writer.WriteString(s)
	ex.writer.WriteString("\n")
}

func (ex *Executor) run(command string) {
	if strings.Contains(command, "push") {
		value := strings.Split(command, " ")[1]
		err := ex.queue.push(value)

		if err != nil {
			ex.writeLn(err.Error())
		}
	}

	if command == "pop" {
		el := ex.queue.pop()
		ex.writeLn(el)
	}

	if command == "peek" {
		el := ex.queue.peek()
		ex.writeLn(el)
	}

	if command == "size" {
		s := ex.queue.size()
		ex.writeLn(strconv.Itoa(s))
	}
}

func createQueue(n int) *Queue {
	arr := make([]string, n, n)
	return &Queue{data: arr, cap: n, head: 0, tail: 0, sz: 0}
}

func (q *Queue) push(x string) error {
	if q.size() == q.cap {
		return errors.New("error")
	}

	q.data[q.tail] = x
	q.tail = (q.tail + 1) % q.cap
	q.sz = q.sz + 1

	return nil
}

func (q *Queue) pop() string {
	if q.size() == 0 {
		return "None"
	}

	el := q.data[q.head]
	q.head = (q.head + 1) % q.cap
	q.sz = q.sz - 1

	return el
}

func (q *Queue) peek() string {
	if q.size() == 0 {
		return "None"
	}

	return q.data[q.head]
}

func (q *Queue) size() int {
	return q.sz
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
