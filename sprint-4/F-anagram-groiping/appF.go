package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var primeCodes map[int32]int

func init() {
	primes := primesEratosthenes(256)
	primeCodes = make(map[int32]int, 256)

	for i, prime := range primes {
		primeCodes[int32(i)] = prime
	}
}

func main() {
	file := openFile("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	solveProblem(reader, writer)
}

func solveProblem(reader *bufio.Reader, writer *bufio.Writer) {
	yaReader := &YaReader{reader}
	_, s := readData(yaReader)

	output := solve(s)

	for _, ot := range output {
		var strmap []string
		for _, o := range ot {
			strmap = append(strmap, strconv.Itoa(o))
		}
		writer.WriteString(strings.Join(strmap, " "))
		writer.WriteString("\n")
	}

	writer.Flush()
}

func solve(words []string) (result [][]int) {
	hashes := make(map[int][]int)

	for i, w := range words {
		hashes[hash(w)] = append(hashes[hash(w)], i)
	}

	for _, indexes := range hashes {
		var entry []int
		for _, index := range indexes {
			entry = append(entry, index)
		}

		result = append(result, entry)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return
}

func hash(s string) (r int) {
	r = 1
	m := 15485863
	for _, ch := range s {
		r = (r % m * primeCodes[ch]) % m
	}

	if r > math.MaxInt32/10 {
		panic(r)
	}

	return
}

func primesEratosthenes(x int) []int {
	n := float64(x)
	size := int(n*math.Log(n) + n*math.Log(math.Log(n)))
	arr := make([]int, size)

	for i := range arr {
		arr[i] = i + 2
	}

	for j := 0; j < size; j++ {
		if arr[j] != 0 {
			for i := j + arr[j]; i < size; i += arr[j] {
				arr[i] = 0
			}
		}

	}

	var r []int
	cnt := 0
	for _, a := range arr {
		if a > 0 {
			r = append(r, a)
			cnt++
			if cnt == x {
				break
			}
		}
	}

	return r
}

func readData(reader *YaReader) (n int, s []string) {
	n = reader.readInt()
	s = strings.Fields(reader.readString())

	return
}

type YaReader struct {
	*bufio.Reader
}

func (reader *YaReader) readString() string {
	line, _ := reader.ReadString('\n')
	return strings.TrimRight(line, "\n")
}

func (reader *YaReader) readInt() int {
	line, _ := reader.ReadString('\n')
	res, _ := strconv.Atoi(strings.TrimRight(line, "\n"))
	return res
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}
