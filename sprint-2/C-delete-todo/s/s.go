package Solution

import "fmt"

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
