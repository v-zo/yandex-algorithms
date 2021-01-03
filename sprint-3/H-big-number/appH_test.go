package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"3\n15 56 2":     "56215",
		"3\n1 783 2":     "78321",
		"3\n1 783 765":   "7837651",
		"5\n2 4 5 2 10":  "542210",
		"4\n9 783 685 7": "97837685",
		"2\n10 1":        "110",
		"3\n57 575 576":  "57657575",
		"3\n831 828 82":  "83182882",
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
