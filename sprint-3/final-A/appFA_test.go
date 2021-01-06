package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"9\n50\n19 21 100 101 1 4 5 7 12":     "-1",
		"9\n5\n19 21 100 101 1 4 5 7 12":      "6",
		"9\n21\n19 21 100 101 500 777 5 7 12": "1",
		"8\n5\n19 100 101 1 4 5 7 12":         "5",
		"8\n1\n19 100 101 1 4 5 7 12":         "3",
		"8\n12\n19 100 101 1 4 5 7 12":        "7",
		"1\n33\n33":                           "0",
		"3\n11\n6 7 11":                       "2",
		"3\n6\n6 7 1":                         "0",
		"3\n7\n6 7 1":                         "1",
		"4\n7\n6 7 1 2":                       "1",
		"4\n1\n6 7 1 2":                       "2",
		"5\n7\n6 7 8 1 2":                     "1",
	}

	for k, v := range cases {
		sr := strings.NewReader(k)
		reader := bufio.NewReader(sr)
		var wr strings.Builder
		writer := bufio.NewWriter(&wr)

		Solve(reader, writer)

		res := strings.Trim(wr.String(), "\n")
		if v != res {
			t.Errorf("\ncase:\n%s\n got: %s\nwant: %s", k, res, v)
		}
	}
}
