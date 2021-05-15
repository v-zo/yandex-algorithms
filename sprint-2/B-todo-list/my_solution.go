package main

import "fmt"

func Solution(head *ListNode) {
	node := head
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}
