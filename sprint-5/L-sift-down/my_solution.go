package main

func siftDown(heap []int, idx int) int {
	left := 2 * idx
	nodesToCompare := []int{idx}
	if left == len(heap) {
		nodesToCompare = append(nodesToCompare, left)
	}
	if left < len(heap) {
		nodesToCompare = append(nodesToCompare, left, left+1)
	}

	if maxIdx := idxOfMax(heap, nodesToCompare...); maxIdx != idx {
		heap[idx-1], heap[maxIdx-1] = heap[maxIdx-1], heap[idx-1]
		maxIdx = siftDown(heap, maxIdx)
		return maxIdx
	}

	return idx
}

func idxOfMax(heap []int, idxs ...int) (maxIdx int) {
	maxIdx = idxs[0]
	for i := 1; i < len(idxs); i++ {
		if heap[idxs[i]-1] > heap[maxIdx-1] {
			maxIdx = idxs[i]
		}
	}

	return
}
