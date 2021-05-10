package main

import "fmt"

func main() {
	node := &Node{1, &Node{2, nil, nil}, &Node{3, nil, nil}}
	fmt.Print(node, "\n")
	remove(node, 3)
	fmt.Print(node)
}
