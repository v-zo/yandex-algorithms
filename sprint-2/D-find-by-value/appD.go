package main

import (
	"fmt"
	"strconv"
)

func main() {
	head := makeNodes(5)

	fmt.Println(Solution(head, "node_4"))
	fmt.Print("\n")

	printNodes(head)
}

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, value string) int {
	node := head
	i := 0
	for node.next != nil {
		if node.data == value {
			return i
		}
		i++
		node = node.next
	}

	return -1
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
