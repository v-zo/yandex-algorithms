package main

import (
	"reflect"
	"testing"
)

type Case struct {
	input    []int
	expected []int
	lf       int
	mid      int
	rg       int
}

func TestMerge(t *testing.T) {
	cases := []*Case{
		{
			[]int{2, 5, 6, 1, 3, 4},
			[]int{1, 2, 3, 4, 5, 6},
			0, 3, 5},
		{
			[]int{3, 5, 6, 7, 1, 2, 4},
			[]int{1, 2, 3, 4, 5, 6, 7},
			0, 4, 6},
		{
			[]int{22, 52, 62, 72, 12, 32, 42},
			[]int{22, 52, 62, 72, 12, 32, 42},
			0, 2, 4},
		{
			[]int{6, 5, 7},
			[]int{5, 6, 7},
			0, 1, 2},
		{
			[]int{4, 3, 5, 0, 1, 2, 3},
			[]int{4, 0, 1, 3, 5, 2, 3},
			1, 3, 4},
	}

	for _, v := range cases {
		res := merge(v.input, v.lf, v.mid, v.rg)
		if !reflect.DeepEqual(v.expected, res) {
			t.Errorf("\ncase:\n%v\n got: %v\nwant: %v", v.input, res, v.expected)
		}
	}
}

type MergeSortCase struct {
	input    *[]int
	expected []int
	lf       int
	rg       int
}

func TestMergeSort(t *testing.T) {
	cases := []*MergeSortCase{
		{
			&[]int{2, 5, 6, 1, 3, 4},
			[]int{1, 2, 3, 4, 5, 6},
			0, 5},
		{
			&[]int{7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7},
			0, 6},
		{
			&[]int{2, 5, 5, 4, 3, 1, 1},
			[]int{1, 1, 2, 3, 4, 5, 5},
			0, 6},
		{
			&[]int{39, 28, 44, 4, 10, 83, 11},
			[]int{4, 10, 11, 28, 39, 44, 83},
			0, 6},
		{
			&[]int{4, 5, 3, 0, 1, 2, 3},
			[]int{4, 0, 1, 3, 5, 2, 3},
			1, 4},
	}

	for _, v := range cases {
		input := append([]int(nil), *v.input...)
		res := append([]int(nil), *v.input...)
		merge_sort(&res, v.lf, v.rg)
		if !reflect.DeepEqual(v.expected, res) {
			t.Errorf("\nmerge-sort:\n%v\n got: %v\nwant: %v", input, res, v.expected)
		}
	}
}
