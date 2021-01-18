package main

import "fmt"

func main() {
	r := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}

	fmt.Println(preorderTraversal(r))
	//fmt.Println(preorderTraversalRecursive(r))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	data []*TreeNode
}

func (stack *Stack) push(el *TreeNode) {
	stack.data = append(stack.data, el)
}

func (stack *Stack) pop() (el *TreeNode, empty bool) {
	el = stack.data[len(stack.data)-1]
	stack.data = stack.data[:len(stack.data)-1]
	empty = len(stack.data) == 0

	return
}

func (stack *Stack) size() int {
	return len(stack.data)
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	stack := &Stack{[]*TreeNode{root}}

	for stack.size() > 0 {
		el, _ := stack.pop()
		res = append(res, el.Val)
		if el.Right != nil {
			stack.push(el.Right)
		}
		if el.Left != nil {
			stack.push(el.Left)
		}
	}

	return res
}

func preorderTraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := preorderTraversalRecursive(root.Left)
	right := preorderTraversalRecursive(root.Right)

	res := []int{root.Val}
	res = append(res, left...)
	res = append(res, right...)

	return res
}
