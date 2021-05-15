package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	expected := 2
	q.Push(expected)
	res := q.Pop()

	if expected != res {
		t.Errorf("\ncase:\n%d\n got: \n%d\nwant: \n%d", expected, res, expected)
	}

	q = NewQueue()
	q.Push(242)
	q.Push(11)

	res = q.Pop()
	if 242 != res {
		t.Errorf("\ncase:\n%d\n got: \n%d\nwant: \n%d", 242, res, 242)
	}

	res = q.Pop()
	if 11 != res {
		t.Errorf("\ncase:\n%d\n got: \n%d\nwant: \n%d", 11, res, 11)
	}
}

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"4 4\n1 2\n2 3\n3 4\n1 4\n3": "3 2 4 1",
		"2 1\n2 1\n1\n":              "1 2",
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
	edges := [][]int{
		{3, 2},
		{5, 4},
		{3, 1},
		{1, 4},
		{1, 6},
		{1, 2},
		{1, 5},
	}
	for i := 0; i < b.N; i++ {
		MainBFS(6, edges, 3)
	}
}
