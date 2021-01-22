package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"10\nget 1\nput 1 10\nput 2 4\nget 1\nget 2\ndelete 2\nget 2\nput 1 5\nget 1\ndelete 2\n":    "None\n10\n4\n4\nNone\n5\nNone",
		"10\nget 1\nput 1 10\nput 2 4\nget 1\nget 2\ndelete 2\ndelete 1\nput 1 5\nget 1\ndelete 2\n": "None\n10\n4\n4\n10\n5\nNone",
		"8\nget 9\ndelete 9\nput 9 1\nget 9\nput 9 2\nget 9\nput 9 3\nget 9\n":                       "None\nNone\n1\n2\n3\n",
		"1\nget 9\n":    "None\n",
		"1\nput 9 4\n":  "",
		"1\ndelete 4\n": "None",
		"5\nput 1 10\nget 1\nput 1 8\nget 1\nget 2\n":    "10\n8\nNone",
		"5\nput 1 10\nget 1\nput 1 8\ndelete 1\nget 1\n": "10\n8\nNone",
	}

	for k, v := range cases {
		sr := strings.NewReader(k)
		reader := bufio.NewReader(sr)
		var wr strings.Builder
		writer := bufio.NewWriter(&wr)

		Solve(reader, writer)

		res := strings.Trim(wr.String(), "\n")
		if strings.Trim(v, "\n") != res {
			t.Errorf("\ncase:\n%s\n got: \n%s\nwant: \n%s", k, res, v)
		}
	}
}
