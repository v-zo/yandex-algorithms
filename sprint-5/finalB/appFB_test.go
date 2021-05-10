package main

import (
	"reflect"
	"testing"
)

type Case struct {
	input    *Node
	item     int
	expected *Node
}

func TestQuickSort(t *testing.T) {
	cases := []Case{
		{
			input:    &Node{1, nil, &Node{2, &Node{3, nil, nil}, nil}},
			item:     3,
			expected: &Node{1, nil, &Node{2, nil, nil}},
		},
	}

	for _, v := range cases {
		result := remove(v.input, v.item)

		if !reflect.DeepEqual(v.expected, result) {
			t.Errorf("\ncase:\n%v\n got: %v\nwant: %v", v.input, result, v.expected)
		}
	}
}
