package main

func siftDown(heap []int, idx int) int {
	left := 2 * idx
	nodesToCompare := []int{idx}
	lengthMinusOne := len(heap) - 1

	if left < lengthMinusOne {
		nodesToCompare = append(nodesToCompare, left, left+1)
	} else if left == lengthMinusOne {
		nodesToCompare = append(nodesToCompare, left)
	}

	if maxIdx := idxOfMax(heap, nodesToCompare...); maxIdx != idx {
		heap[idx], heap[maxIdx] = heap[maxIdx], heap[idx]
		maxIdx = siftDown(heap, maxIdx)
		return maxIdx
	}

	return idx
}

func idxOfMax(heap []int, idxs ...int) (maxIdx int) {
	maxIdx = idxs[0]
	for i := 1; i < len(idxs); i++ {
		if heap[idxs[i]] > heap[maxIdx] {
			maxIdx = idxs[i]
		}
	}

	return
}
