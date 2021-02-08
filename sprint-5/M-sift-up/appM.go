package main

import "fmt"

func main() {
	heap := []int{10, 6, 2, 8}
	answer := siftUp(heap, 3)
	fmt.Println(answer)
}
