package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func v1(w io.Writer, s string) {
	io.WriteString(w, s+"\n")
}

func v2(w io.Writer, s string) {
	io.WriteString(w, s)
	io.WriteString(w, "\n")
}

func Benchmark_v1(b *testing.B) {
	w := bufio.NewWriter(ioutil.Discard)
	short := strings.Split("some", " ")[0] // split needed for avoiding const optimizations
	long := strings.Split("some_string_very_long_nonsence_bla_bla_bla_bla_bla_string_long_nonsnce", " ")[0]

	b.Run("short", func(b *testing.B) {
		b.Run("concat", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				v1(w, short)
			}
		})
		b.Run("two calls", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				v2(w, short)
			}
		})
	})
	b.Run("long", func(b *testing.B) {
		b.Run("concat", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				v1(w, long)
			}
		})
		b.Run("two calls", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				v2(w, long)
			}
		})
	})
}
