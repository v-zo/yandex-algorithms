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
		"6 7\n3 2\n5 4\n3 1\n1 4\n1 6\n1 2\n1 5\n1": "1 2 3 4 5 6",
		"7 6\n3 6\n2 7\n4 5\n6 5\n2 4\n1 3\n4":      "4 2 7 5 6 3 1",
		"7 6\n7 4\n7 3\n6 7\n7 2\n5 7\n7 1\n1":      "1 7 2 3 4 5 6",
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
