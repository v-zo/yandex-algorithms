package main

import "fmt"

func main() {
	heap := []int{0, 10, 6, 2, 8}
	answer := siftDown(heap, 3)
	fmt.Println(answer)
}
