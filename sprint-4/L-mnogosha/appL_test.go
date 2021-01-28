package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"10 2\ngggggooooogggggoooooogggggssshaa\n": "0 5",
		"3 4\nallallallallalla\n":                  "0 1 2",
		"1 1\nx\n":                                 "0",
	}

	for k, v := range cases {
		sr := strings.NewReader(k)
		reader := bufio.NewReader(sr)
		var wr strings.Builder
		writer := bufio.NewWriter(&wr)

		process(reader, writer)

		err := writer.Flush()
		check(err)

		res := strings.Trim(wr.String(), "\n")
		if v != res {
			t.Errorf("\n- case:\n%s- got: \n%s\n- want: \n%s", k, res, v)
		}
	}
}
