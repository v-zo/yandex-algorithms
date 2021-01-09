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

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go writeBenchData(fibonacciBigInt, "data.txt", &wg)
	go writeBenchData(fibonacciMatrix, "dataM.txt", &wg)
	wg.Wait()
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

func newBigInt(n int) *big.Int {
	return new(big.Int).SetUint64(uint64(n))
}

func findFib(n int, p, p0 *big.Int) *big.Int {
	for i := 0; i < n; i++ {
		p, p0 = p0, p
		p.Add(p, p0)
	}

	return p
}

func fibonacciBigInt(n int) *big.Int {
	return findFib(n, newBigInt(1), newBigInt(0))
}

type MatrixBigInt struct {
	value [][]*big.Int
}

func NewMatrixBigInt(n int) MatrixBigInt {
	M := make([][]*big.Int, n)
	for i := range M {
		M[i] = make([]*big.Int, n)
		for j := range M {
			M[i][j] = newBigInt(0)
		}
	}

	return MatrixBigInt{M}
}

func FromInt(A [][]int) [][]*big.Int {
	n := len(A)
	M := make([][]*big.Int, n)
	for i := range M {
		M[i] = make([]*big.Int, n)
		for j := range A {
			M[i][j] = newBigInt(A[i][j])
		}
	}

	return M
}

func CreateUnitMatrix(size int) [][]*big.Int {
	M := NewMatrixBigInt(size)

	for i := 0; i < size; i++ {
		M.value[i][i] = newBigInt(1)
	}

	return M.value
}

func product(A, B [][]*big.Int) [][]*big.Int {
	n := len(A)

	prod := NewMatrixBigInt(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				mult := newBigInt(n)
				mult.Mul(A[i][k], B[k][j])
				prod.value[i][j] = prod.value[i][j].Add(prod.value[i][j], mult)
			}
		}
	}

	return prod.value
}

func matrixPower(A [][]*big.Int, n int) [][]*big.Int {
	if n == 1 {
		return A
	}

	p := CreateUnitMatrix(len(A))

	if n == 0 {
		return p
	}

	for n != 0 {
		if n%2 != 0 {
			p = product(p, A)
		}

		A = product(A, A)
		n /= 2
	}

	return p
}

func fibonacciMatrix(n int) *big.Int {
	if n == 0 || n == 1 {
		return newBigInt(1)
	}

	F := FromInt([][]int{{1, 1}, {1, 0}})

	return matrixPower(F, n)[0][0]
}

func createFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}
