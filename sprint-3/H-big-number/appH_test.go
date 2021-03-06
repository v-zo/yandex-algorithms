package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"3\n15 56 2":          "56215",
		"3\n1 783 2":          "78321",
		"3\n1 783 765":        "7837651",
		"5\n2 4 5 2 10":       "542210",
		"4\n9 783 685 7":      "97837685",
		"2\n10 1":             "110",
		"3\n57 575 576":       "57657575",
		"3\n831 828 82":       "83182882",
		"6\n388 86 8 83 82 3": "88683823883",
		"6\n88 86 8 83 82 3":  "8888683823",
		"2\n10 01":            "1001",
		"2\n10 11":            "1110",
		"2\n309 33":           "33309",
		"3\n309 33 3":         "333309",
		"3\n309 33 03":        "3330903",
		"0\n":                 "",
		"1\n2":                "2",
		"1\n0":                "0",
		"3\n1 10 101":         "110110",
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
