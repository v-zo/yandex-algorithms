package main

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"3\ni love coffee\ncoffee with milk and sugar\nfree tea for everyone\n3\ni like black coffee without milk\neveryone loves new year\nmary likes black coffee without milk\n":                          "1 2\n3\n2 1",
		"6\nbuy flat in moscow\nrent flat in moscow\nsell flat in moscow\nwant flat in moscow like crazy\nclean flat in moscow on weekends\nrenovate flat in moscow\n1\nflat in moscow for crazy weekends\n": "4 5 1 2 3",
		"3\na a a a\nb b a a\nb b a a\n3\na b b\nx\nz\n": "1 2 3",
	}

	for k, v := range cases {
		sr := strings.NewReader(k)
		reader := bufio.NewReader(sr)
		var wr strings.Builder
		writer := bufio.NewWriter(&wr)

		Solve(reader, writer)

		writer.Flush()

		res := strings.Trim(wr.String(), "\n")
		if v != res {
			t.Errorf("\ncase:\n%s\n got: \n%s\nwant: \n%s", k, res, v)
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	dmax := 50
	wmax := 500
	rmax := 1
	docs := strconv.Itoa(dmax) + "\n"
	for i := 0; i < dmax; i++ {
		docs += "a"
		for j := 0; j < wmax; j++ {
			docs += " a"
		}
		docs += "\n"
	}
	docs += strconv.Itoa(rmax) + "\n"
	for i := 0; i < rmax; i++ {
		docs += "b"
		for j := 0; j < wmax; j++ {
			docs += " b"
		}
		docs += "\n"
	}
	docs += "\n"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sr := strings.NewReader(docs)
		reader := bufio.NewReader(sr)
		var wr strings.Builder
		writer := bufio.NewWriter(&wr)

		Solve(reader, writer)
	}
}
