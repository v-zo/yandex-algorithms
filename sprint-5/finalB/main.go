package main

import "fmt"

func main() {
	node := &Node{1, &Node{2, nil, nil}, &Node{4, nil, nil}}
	remove(node, 2)
	fmt.Print(node, "\n")

	node = &Node{5,
		&Node{3,
			&Node{2, nil, nil},
			&Node{4, nil, nil},
		},
		&Node{6, nil, &Node{7, nil, nil}},
	}

	remove(node, 3)
	fmt.Print(node, "\n")
}
