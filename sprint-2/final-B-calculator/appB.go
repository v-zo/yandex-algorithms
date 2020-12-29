/*
-- ПРИНЦИП РАБОТЫ --


-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

*/

package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	file := openFile("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	Solve(reader, writer)
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	ch := scanner.Text()
	stack := &Stack{&Node{ch, nil}}

	for scanner.Scan() {
		//writer.WriteString(scanner.Text())
		ch = scanner.Text()

		isOperator := strings.Contains("+-*/", ch)

		if isOperator {

		} else {
			stack.push(ch)
		}

		//writer.WriteString(stack.head.value)
	}

	//writer.WriteString("\n")

	writer.Flush()
}

type Node struct {
	value string
	prev  *Node
}

type Stack struct {
	head *Node
}

func (st *Stack) push(val string) {
	st.head = &Node{val, st.head}
}

func (st *Stack) pop() string {
	st.head = &Node{st.head.value, st.head.prev}

	return st.head.value
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
