package main

import "fmt"

func main() {
	tree := &TNode{5,
		&TNode{3, &TNode{8, nil, nil}, nil},
		&TNode{7, nil, nil},
	}
	max := Solution(tree)
	fmt.Println(max)
}

type TNode struct {
	value int
	left  *TNode
	right *TNode
}
