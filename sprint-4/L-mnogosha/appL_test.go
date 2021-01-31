package main

import (
	"bufio"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"10 2\ngggggooooogggggoooooogggggssshaa\n": "0 5",
		"3 4\nallallallallalla\n":                  "0 1 2",
		"1 1\nx\n":                                 "0",
		"1 2\nxx\n":                                "0",
		"1 1\nxz\n":                                "0 1",
	}

	for k, v := range cases {
		sr := strings.NewReader(k)
		reader := bufio.NewReader(sr)
		var wr strings.Builder
		writer := bufio.NewWriter(&wr)

		processData(reader, writer)

		err := writer.Flush()
		check(err)

		res := strings.Trim(wr.String(), "\n")

		if !reflect.DeepEqual(splitNSort(res), splitNSort(v)) {
			t.Errorf("\n- case:\n%s- got: \n%s\n- want: \n%s", k, res, v)
		}
	}
}

func splitNSort(s string) []string {
	arr := strings.Split(s, " ")
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return arr
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(3, 4, "allallallallalla")
	}
}
