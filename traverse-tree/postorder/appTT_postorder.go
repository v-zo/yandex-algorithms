package main

import "fmt"

func main() {
	r := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}

	//fmt.Println(postorderTraversal(r))
	fmt.Println(postorderTraversalRecursive(r))
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

func postorderTraversal(root *TreeNode) []int {
	//stack:=&Stack{}
	//curr := root
	var rs []int

	// TODO

	return rs
}

func postorderTraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	res = append(res, postorderTraversalRecursive(root.Left)...)
	res = append(res, postorderTraversalRecursive(root.Right)...)
	res = append(res, root.Val)
	return res
}
