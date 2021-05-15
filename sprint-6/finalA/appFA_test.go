package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"4 4\n1 2 5\n1 3 6\n2 4 8\n3 4 3\n":        "19",
		"3 3\n1 2 1\n1 2 2\n2 3 1\n":               "3",
		"2 0\n":                                    "Oops! I did it again",
		"5 5\n1 2 1\n1 2 2\n2 3 1\n4 5 1\n5 4 2\n": "Oops! I did it again",
		"1 0\n": "0",
		"10 20\n9 10 4\n2 2 4\n4 2 8\n10 5 3\n1 10 6\n7 4 2\n10 10 6\n3 7 4\n8 9 4\n8 10 7\n6 10 10\n2 8 8\n3 8 1\n3 10 3\n9 5 8\n10 10 2\n1 8 1\n10 1 5\n3 6 10\n9 10 8\n": "69",
		"10 20\n8 7 9\n4 10 7\n6 6 2\n8 10 1\n10 6 1\n5 7 8\n1 9 6\n10 3 3\n10 5 8\n6 6 6\n5 7 9\n5 2 4\n3 1 1\n10 7 8\n8 4 6\n5 5 7\n7 8 6\n5 10 2\n10 1 3\n3 5 9\n":       "56",
		"10 20\n8 10 10\n1 7 1\n8 7 5\n2 10 9\n3 5 5\n5 5 6\n4 3 2\n4 5 8\n8 5 5\n4 7 6\n6 2 10\n9 7 7\n3 10 10\n3 8 7\n6 2 5\n8 5 3\n6 1 7\n6 8 7\n9 1 7\n1 1 4\n":         "74",
		"10 20\n2 8 6\n7 2 1\n2 9 4\n5 4 9\n8 9 1\n7 10 3\n10 10 8\n9 2 5\n4 3 5\n1 10 5\n6 2 10\n2 6 10\n8 2 8\n9 2 3\n10 2 1\n3 10 3\n2 8 5\n2 8 4\n7 1 1\n7 5 4\n":       "50",
		"10 20\n5 1 7\n8 7 10\n8 8 9\n10 6 2\n3 3 3\n10 4 8\n9 10 2\n2 3 3\n10 5 1\n7 3 8\n2 5 8\n2 3 4\n10 8 7\n5 6 8\n1 6 2\n2 4 6\n10 7 2\n8 3 1\n10 8 5\n5 9 6\n":       "68",
		"10 20\n1 7 8\n3 3 6\n10 2 4\n6 8 10\n6 3 9\n3 9 10\n6 1 10\n9 2 7\n2 3 3\n1 9 7\n2 5 2\n1 6 2\n3 6 5\n5 4 2\n6 5 4\n8 10 1\n6 4 6\n1 5 5\n3 6 8\n6 9 3\n":          "69",
		"10 20\n9 8 3\n3 1 5\n10 4 9\n4 1 4\n9 7 6\n6 8 8\n8 8 10\n5 2 2\n3 8 1\n5 6 8\n5 5 7\n10 10 7\n7 9 6\n4 3 7\n6 2 6\n6 6 6\n8 2 5\n2 10 1\n5 3 8\n7 9 6\n":          "60",
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

func BenchmarkSample(b *testing.B) {
	file := openFile("input.txt")
	defer file.close()
	n, edges := readData(bufio.NewReader(file))

	graph := NewGraph(n, edges)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		findMST(graph, n)
	}
}
