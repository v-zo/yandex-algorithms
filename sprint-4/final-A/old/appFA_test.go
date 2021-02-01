package main

import (
	"bufio"
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

//type Case struct {
//	input    []Rel
//	expected
//}
//
//func TestSortRelSlice(t *testing.T) {
//	cases := map[Case]string{
//	}
//
//
//
//	for k, v := range cases {
//		res := sortRelSlice()
//
//		if v != res {
//			t.Errorf("\ncase:\n%s\n got: \n%s\nwant: \n%s", k, res, v)
//		}
//	}
//}
