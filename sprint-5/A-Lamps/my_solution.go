package main

import (
	"math"
)

func Solution(root *TNode) (max int) {
	stack := &Stack{[]*TNode{root}}
	max = math.MinInt64

	for stack.size() > 0 {
		el, _ := stack.pop()

		if el.value > max {
			max = el.value
		}

		if el.right != nil {
			stack.push(el.right)
		}
		if el.left != nil {
			stack.push(el.left)
		}
	}

	return
}

type Stack struct {
	data []*TNode
}

func (stack *Stack) push(el *TNode) {
	stack.data = append(stack.data, el)
}

func (stack *Stack) pop() (el *TNode, empty bool) {
	el = stack.data[len(stack.data)-1]
	stack.data = stack.data[:len(stack.data)-1]
	empty = len(stack.data) == 0

	return
}

func (stack *Stack) size() int {
	return len(stack.data)
}
