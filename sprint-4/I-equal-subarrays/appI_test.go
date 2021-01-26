package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"5\n1 2 3 2 1\n5\n3 2 1 5 6\n": "3",
		"5\n1 2 3 4 5\n3\n4 5 9\n":     "2",
		"5\n1 1 1 2 2\n4\n1 1 2 2 2\n": "4",
	}

	for k, v := range cases {
		sr := strings.NewReader(k)
		reader := bufio.NewReader(sr)
		var wr strings.Builder
		writer := bufio.NewWriter(&wr)

		process(reader, writer)

		writer.Flush()

		res := strings.Trim(wr.String(), "\n")
		if v != res {
			t.Errorf("\n- case:\n%s- got: \n%s\n- want: \n%s", k, res, v)
		}
	}
}
