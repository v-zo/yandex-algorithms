package main

import "fmt"

func main() {
	r := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}

	fmt.Println(inorderTraversal(r))
	//fmt.Println(inorderTraversalRecursive(r))
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

func inorderTraversal(root *TreeNode) []int {
	stack := &Stack{}
	curr := root
	var res []int
	for curr != nil || stack.size() > 0 {
		for curr != nil {
			stack.push(curr)
			curr = curr.Left
		}
		curr, _ = stack.pop()
		res = append(res, curr.Val)
		curr = curr.Right
	}

	return res
}

func inorderTraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int

	if root.Left != nil {
		res = append(res, inorderTraversalRecursive(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, inorderTraversalRecursive(root.Right)...)
	}

	return res
}
