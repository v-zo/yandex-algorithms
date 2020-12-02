package main

import (
	"reflect"
	"testing"
)

func TestGetZeroesPositions(t *testing.T) {
	got := GetZeroesPositions([]string{"0", "1", "4", "9", "0"})
	if !reflect.DeepEqual(got, []int{0, 4}) {
		t.Errorf("GetZeroesPositions([]string{\"0\", \"1\", \"4\", \"9\", \"0\"}) = %d; want []int{0, 4}", got)
	}
}

func TestMakeRangeReversed(t *testing.T) {
	got := MakeRangeReversed(20, 10)
	if !reflect.DeepEqual(got, []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10}) {
		t.Errorf("MakeRangeReversed(20, 10) = %d; want [20 19 18 17 16 15 14 13 12 11 10]", got)
	}
}
