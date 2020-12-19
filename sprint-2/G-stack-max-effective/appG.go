package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	data []int
	max  int
}

func main() {
	file := openFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	stack := createStack()

	for scanner.Scan() {
		executeCommand(stack, scanner.Text())
	}
}

func createStack() *Stack {
	return &Stack{[]int{}, math.MinInt64}
}

func (stack *Stack) top() int {
	return stack.data[len(stack.data)-1]
}

func (stack *Stack) push(x int) {
	val := x
	if x > stack.max {
		val = 2*x - stack.max
		stack.max = x
	}

	stack.data = append(stack.data, val)
}

func (stack *Stack) pop() string {
	if len(stack.data) == 0 {
		return "error"
	}

	val := stack.top()
	if val > stack.max {
		stack.max = 2*stack.max - val
	}

	stack.data = stack.data[:len(stack.data)-1]
	return ""
}

func (stack *Stack) getMax() string {
	if len(stack.data) == 0 {
		return "None"
	}

	return strconv.Itoa(stack.max)
}

func executeCommand(stack *Stack, line string) {
	if line == "get_max" {
		fmt.Println(stack.getMax())
	}

	if line == "pop" {
		err := stack.pop()
		if err == "error" {
			fmt.Println("error")
		}
	}

	res := strings.Split(line, " ")
	if res[0] == "push" {
		num, _ := strconv.Atoi(res[1])
		stack.push(num)
	}
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
