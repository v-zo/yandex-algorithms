/*
-- ПРИНЦИП РАБОТЫ --


-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

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

type Node struct {
	value int
	prev  *Node
}

type Stack struct {
	head *Node
}

type Calculator struct {
	stack *Stack
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	ch := scanner.Text()
	calculator := createCalculator(ch)

	for scanner.Scan() {
		ch = scanner.Text()
		calculator.process(ch)
	}

	writer.WriteString(strconv.Itoa(calculator.getResult()))
	writer.WriteString("\n")
	writer.Flush()
}

func createCalculator(initialChar string) *Calculator {
	num, _ := strconv.Atoi(initialChar)
	stack := &Stack{&Node{num, nil}}

	return &Calculator{stack}
}

func (clr *Calculator) getResult() int {
	return clr.stack.head.value
}

func (clr *Calculator) process(ch string) {
	isOperator := strings.Contains("+-*/", ch)

	if isOperator {
		switch ch {
		case "+":
			a := clr.stack.pop()
			b := clr.stack.pop()
			clr.stack.push(a + b)
		case "-":
			a := clr.stack.pop()
			b := clr.stack.pop()
			clr.stack.push(b - a)
		case "*":
			a := clr.stack.pop()
			b := clr.stack.pop()
			clr.stack.push(a * b)
		case "/":
			a := clr.stack.pop()
			b := clr.stack.pop()
			clr.stack.push(dev(b, a))
		}
	} else {
		num, _ := strconv.Atoi(ch)
		clr.stack.push(num)
	}
}

func dev(a int, b int) int {
	abs := func(i int) int {
		if i < 0 {
			return -i
		}

		return i
	}

	if abs(a) < abs(b) && a < 0 && b > 0 {
		return a
	} else {
		return a / b
	}
}

func (stack *Stack) push(val int) {
	stack.head = &Node{val, stack.head}
}

func (stack *Stack) pop() (value int) {
	value = stack.head.value
	stack.head = stack.head.prev

	return
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
