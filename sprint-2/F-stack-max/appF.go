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
}

func main() {
	file := openFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	stack := &Stack{}

	for scanner.Scan() {
		executeCommand(stack, scanner.Text())
	}
}

func (stack *Stack) push(x int) {
	stack.data = append(stack.data, x)
}

func (stack *Stack) pop() string {
	if len(stack.data) == 0 {
		return "error"
	}

	stack.data = stack.data[:len(stack.data)-1]
	return ""
}

func (stack *Stack) getMax() string {
	if len(stack.data) == 0 {
		return "None"
	}

	return strconv.Itoa(maxInt(stack.data))
}

func executeCommand(stack *Stack, line string) {
	if line == "get_max" {
		fmt.Println(stack.getMax())
	}

	if line == "pop" {
		error := stack.pop()
		if error == "error" {
			fmt.Println("error")
		}
	}

	res := strings.Split(line, " ")
	if res[0] == "push" {
		num, _ := strconv.Atoi(res[1])
		stack.push(num)
	}
}

func maxInt(arr []int) int {
	m := math.MinInt64
	for _, v := range arr {
		if v > m {
			m = v
		}
	}

	return m
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
