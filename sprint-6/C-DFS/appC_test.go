package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"4 4\n3 2\n4 3\n1 4\n1 2\n3": "3 2 1 4",
		"2 1\n1 2\n1":                "1 2",
		"3 1\n2 3\n1":                "1",
		"6 7\n3 2\n5 4\n3 1\n1 4\n1 6\n1 2\n1 5\n1":                                      "1 2 3 4 5 6",
		"7 6\n3 6\n2 7\n4 5\n6 5\n2 4\n1 3\n4":                                           "4 2 7 5 6 3 1",
		"7 6\n7 4\n7 3\n6 7\n7 2\n5 7\n7 1\n1":                                           "1 7 2 3 4 5 6",
		"6 7\n4 5\n5 1\n1 6\n5 2\n4 2\n1 4\n4 3\n2\n":                                    "2 4 1 5 6 3",
		"6 7\n1 3\n3 5\n4 6\n6 2\n1 5\n1 6\n1 2\n6\n":                                    "6 1 2 3 5 4",
		"10 13\n9 2\n1 3\n6 10\n5 10\n9 10\n5 9\n3 5\n5 2\n6 7\n5 6\n8 3\n6 1\n4 2\n7\n": "7 6 1 3 5 2 4 9 10 8",
		"10 11\n10 9\n5 10\n7 1\n3 8\n2 9\n9 6\n1 5\n3 1\n4 1\n5 3\n3 7\n5\n":            "5 1 3 7 8 4 10 9 2 6",
		"10 12\n6 3\n2 10\n3 1\n5 3\n6 7\n8 7\n1 2\n2 6\n10 5\n9 1\n10 7\n3 4\n5\n":      "5 3 1 2 6 7 8 10 9 4",
	}

	for k, v := range cases {
		sr := strings.NewReader(k)
		reader := bufio.NewReader(sr)
		var wr strings.Builder
		writer := bufio.NewWriter(&wr)

		Solve(reader, writer)

		err := writer.Flush()
		check(err)

		res := strings.Trim(wr.String(), "\n")
		if strings.Trim(v, "\n") != res {
			t.Errorf("\ncase:\n%s\n got: \n%s\nwant: \n%s", k, res, v)
		}
	}
}

func BenchmarkSample(b *testing.B) {
	edges := [][]int{
		{3, 2},
		{5, 4},
		{3, 1},
		{1, 4},
		{1, 6},
		{1, 2},
		{1, 5},
	}
	for i := 0; i < b.N; i++ {
		MainDFS(6, edges, 3)
	}
}
