package main

func siftUp(heap []int, idx int) int {
	if idx == 1 {
		return idx
	}

	parentIdx := idx / 2

	if heap[parentIdx] < heap[idx] {
		heap[idx], heap[parentIdx] = heap[parentIdx], heap[idx]
		newIdx := siftUp(heap, parentIdx)
		return newIdx
	}

	return idx
}
