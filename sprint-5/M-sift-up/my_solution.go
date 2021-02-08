package main

func siftUp(heap []int, idx int) int {
	if idx == 1 {
		return idx
	}

	parentIdx := idx / 2

	if heap[parentIdx-1] < heap[idx-1] {
		heap[idx-1], heap[parentIdx-1] = heap[parentIdx-1], heap[idx-1]
		newIdx := siftUp(heap, parentIdx)
		return newIdx
	}

	return idx
}
