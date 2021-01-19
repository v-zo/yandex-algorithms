package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"1000\n1000000009\nabcdefgh\n7\n1 1\n1 5\n2 3\n3 4\n4 4\n1 8\n5 8": "97\n98226219\n98099\n99100\n100\n218067142\n102102195",
		"100\n10\na\n1\n1 1\n": "7",
	}

	for k, v := range cases {
		sr := strings.NewReader(k)
		reader := bufio.NewReader(sr)
		var wr strings.Builder
		writer := bufio.NewWriter(&wr)

		Solve(reader, writer)

		res := strings.Trim(wr.String(), "\n")
		if v != res {
			t.Errorf("\ncase:\n%s\n got: \n%s\nwant: \n%s", k, res, v)
		}
	}
}
