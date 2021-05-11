package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"4 4\n1 2 5\n1 3 6\n2 4 8\n3 4 3\n": "19",
		"3 3\n1 2 1\n1 2 2\n2 3 1\n":        "3",
		"2 0\n":                             "Oops! I did it again",
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
