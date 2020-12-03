package main

import (
	"reflect"
	"testing"
)

func TestGetPrimes(t *testing.T) {
	got := GetPrimes(100)
	if !reflect.DeepEqual(got, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}) {
		t.Errorf(" = %d; want [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97]", got)
	}
}
