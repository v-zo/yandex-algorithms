package main

import (
	"reflect"
	"testing"
)

func TestHorner(t *testing.T) {
	cases := map[string][]int{
		"abcdefgh": {97, 97098, 97098099, 98098227, 98226219, 226218220, 218218069, 218067142},
	}

	for s, v := range cases {

		res := horner(1000, 1000000009, s)

		//fmt.Println(res)

		if !reflect.DeepEqual(res, v) {
			t.Errorf("\ncase:\n%s\n got: \n%v\nwant: \n%v", s, res, v)
		}
	}
}
