package main

import (
	"fmt"
	"strconv"
)

func main() {
	head := makeNodes(7)

	newHead := Solution(head)

	printNodes(newHead)
}

type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}

func Solution(head *ListNode) *ListNode {
	node := head
	nextHead := &ListNode{node.data, nil, nil}
	newHead := &ListNode{node.next.data, nextHead, nil}
	node = node.next.next
	nextHead.prev = newHead

	for node != nil {
		nextHead = newHead
		newHead = &ListNode{node.data, nextHead, nil}
		nextHead.prev = newHead
		node = node.next
	}

	return newHead
}

func printNodes(head *ListNode) {
	node := head
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}

func makeNodes(max int) *ListNode {
	next := &ListNode{"node_" + strconv.Itoa(max), nil, nil}
	node := &ListNode{"node_" + strconv.Itoa(max-1), next, nil}
	next.prev = node
	for i := max - 2; i > 0; i-- {
		next = node

		node = &ListNode{"node_" + strconv.Itoa(i), next, nil}
		next.prev = node
	}

	return node
}
