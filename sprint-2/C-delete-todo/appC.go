package main

import (
	"fmt"
	"strconv"
)

func main() {
	head := makeNodes(5)

	newHead := Solution(head, 1)

	printNodes(newHead)
}

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, index int) *ListNode {
	if index == 0 {
		return head.next
	}

	node := head
	var prev *ListNode

	for index > 0 {
		index -= 1
		prev = node
		node = node.next
	}

	prev.next = node.next

	return head
}

func printNodes(head *ListNode) {
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
