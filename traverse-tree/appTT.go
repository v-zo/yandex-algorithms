package main

import "fmt"

func main() {
	r := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}

	//fmt.Println(preorderTraversal(r))
	fmt.Println(preorderTraversalRecursive(r))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func preorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return []int{}
//	}
//
//	res := []int{root.Val}
//
//	node := root.Left
//	for node != nil {
//		res=append(res, node.Val)
//		node = node.Left
//	}
//
//	node = root.Right
//	for node != nil {
//		res=append(res, node.Val)
//		node = node.Right
//	}
//
//	return res
//}

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
