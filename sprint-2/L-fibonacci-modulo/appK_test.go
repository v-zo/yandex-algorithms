package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestApp(t *testing.T) {
	cases := map[string]int{
		"3 1":  3,
		"10 1": 9,
		"1 1":  1,
		"0 1":  1,
		"10 5": 89,
	}

	for input, expected := range cases {
		fields := strings.Fields(input)
		n, _ := strconv.Atoi(fields[0])
		m, _ := strconv.Atoi(fields[1])
		res := fibonacciModulo(n, m, 1, 0, 0)
		if expected != res {
			t.Errorf("\ncase:\n%s\n got: %d\nwant: %d", input, res, expected)
		}
	}
}
