package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"sync"
	"testing"
	"text/template"
)

const (
	from = 200
	to   = 2200
	dots = 200
)

type Entry struct {
	X string
	Y string
}

func writeBenchData(fn func(int) *big.Int, fileName string, wg *sync.WaitGroup) {
	defer wg.Done()

	const entry = "{{`{`}}{{ .X }},{{ .Y  }}{{`},`}}"
	t := template.Must(template.New("entry").Parse(entry))

	file := createFile(fileName)
	defer file.Close()
	w := bufio.NewWriter(file)

	delta := (to - from) / dots

	for n := from; n < to; n += delta {
		fmt.Printf("\033[2K\r%s%d...", "running bench, n=", n)

		res := testing.Benchmark(func(b *testing.B) {
			BenchFib(b.N, n, fn)
		})

		y := strconv.Itoa(int(res.NsPerOp()))
		x := strconv.Itoa(n)
		point := Entry{x, y}
		t.Execute(w, point)
	}

	fmt.Printf("\033[2K\r")
	w.Flush()
	fmt.Printf("Done.\n")
}

func BenchFib(N, p int, fn func(int) *big.Int) {
	for i := 0; i < N; i++ {
		fn(p)
	}
}

func createFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}
