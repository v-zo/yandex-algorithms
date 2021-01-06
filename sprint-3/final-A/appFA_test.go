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
		"1\n33\n33":                           "0",
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
