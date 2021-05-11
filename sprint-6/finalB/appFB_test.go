package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"3\nRB\nR\n":            "NO",
		"4\nBBB\nRB\nB\n":       "YES",
		"5\nRRRB\nBRR\nBR\nR\n": "NO",
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
