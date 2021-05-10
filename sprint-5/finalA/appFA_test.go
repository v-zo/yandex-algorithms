package main

import (
	"reflect"
	"testing"
)

type SortableInt []int

type Case struct {
	input    SortableInt
	expected SortableInt
}

func (c SortableInt) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c SortableInt) Swap(i int, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c SortableInt) Len() int {
	return len(c)
}

func TestQuickSort(t *testing.T) {
	cases := []Case{
		{input: SortableInt{3, 2, 1}, expected: SortableInt{1, 2, 3}},
		{input: SortableInt{4, 3, 2, 1}, expected: SortableInt{1, 2, 3, 4}},
		{input: SortableInt{4, 3, 5, 1}, expected: SortableInt{1, 3, 4, 5}},
		{input: SortableInt{3, 2, 4, 1}, expected: SortableInt{1, 2, 3, 4}},
		{input: SortableInt{3, 5, 2, 4, 1}, expected: SortableInt{1, 2, 3, 4, 5}},
		{input: SortableInt{3, 5, 10, 4, 1}, expected: SortableInt{1, 3, 4, 5, 10}},
		{input: SortableInt{3, 5, 0, 4, 1}, expected: SortableInt{0, 1, 3, 4, 5}},
		{input: SortableInt{}, expected: SortableInt{}},
		{input: SortableInt{0}, expected: SortableInt{0}},
		{input: SortableInt{0, 1}, expected: SortableInt{0, 1}},
		{input: SortableInt{1, 0}, expected: SortableInt{0, 1}},
		{input: SortableInt{4, 8, 9, 20, 1, 5, 3, 10}, expected: SortableInt{1, 3, 4, 5, 8, 9, 10, 20}},
	}

	for _, v := range cases {
		cp := append(SortableInt{}, v.input...)
		heapSort(cp)

		if !reflect.DeepEqual(v.expected, cp) {
			t.Errorf("\ncase:\n%v\n got: %v\nwant: %v", v.input, cp, v.expected)
		}
	}
}
