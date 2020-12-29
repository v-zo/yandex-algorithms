/*
посылка 45999658

-- ПРИНЦИП РАБОТЫ --
Весь принцип работы уже расписан в задании, осталось только его реализовать.
Используем стек на основе связного списка

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Алгоритм точно следует инструкции, описанной в задании.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Сложность команд стека (pop, push, top): 0(1), т.к. все они сводятся к операциям добавления/удаления элементов в связном списке, а это O(1)
Сложность операций над операндами:  O(1), т.к. по сути сводятся к простым арифметическим оперцаиям + командам стека

В итоге: O(1)
*/

package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
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
	return clr.stack.top()
}

func (clr *Calculator) process(ch string) {
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
		clr.stack.push(floorDiv(b, a))
	default:
		num, _ := strconv.Atoi(ch)
		clr.stack.push(num)
	}
}

func floorDiv(a int, b int) int {
	if a*b < 0 {
		fa := float64(a)
		fb := float64(b)

		return int(math.Floor(fa / fb))
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

func (stack *Stack) top() int {
	return stack.head.value
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
