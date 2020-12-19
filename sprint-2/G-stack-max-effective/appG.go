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

type EffectiveStack struct {
	stack Stack
	max   Stack
}

func main() {
	file := openFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	startStack := Stack{[]int{}}
	maxStack := Stack{[]int{math.MinInt64}}
	es := &EffectiveStack{startStack, maxStack}

	for scanner.Scan() {
		executeCommand(es, scanner.Text())
	}
}

func (stack *Stack) peek() int {
	return stack.data[len(stack.data)-1]
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

func (es *EffectiveStack) push(x int) {
	es.stack.data = append(es.stack.data, x)

	if x > es.max.peek() {
		es.max.push(x)
	}
}

func (es *EffectiveStack) pop() string {
	if len(es.stack.data) == 0 {
		return "error"
	}

	lastEl := es.stack.peek()
	max := es.max.peek()

	if lastEl == max {
		es.max.pop()
	}

	es.stack.data = es.stack.data[:len(es.stack.data)-1]
	return ""
}

func (es *EffectiveStack) getMax() string {
	if len(es.stack.data) == 0 {
		return "None"
	}

	return strconv.Itoa(es.max.peek())
}

func executeCommand(es *EffectiveStack, line string) {
	if line == "get_max" {
		fmt.Println(es.getMax())
	}

	if line == "pop" {
		error := es.pop()
		if error == "error" {
			fmt.Println("error")
		}
	}

	res := strings.Split(line, " ")
	if res[0] == "push" {
		num, _ := strconv.Atoi(res[1])
		es.push(num)
	}
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
