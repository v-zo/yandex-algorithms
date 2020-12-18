package main

import (
	"fmt"
	"strconv"
)

func main() {
	head := makeNodes(5)

	Solution(head)
}

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode) {
	node := head
	for node.next != nil {
		fmt.Println(node.data)
		node = node.next
	}
}

func makeNodes(max int) *ListNode {
	node := &ListNode{"node_" + strconv.Itoa(max), nil}
	for i := max - 1; i >= 0; i-- {
		node = &ListNode{"node_" + strconv.Itoa(i), node}
	}

	return node
}
