package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file := openFile("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	Solve(reader, writer)
}

type KV struct {
	key string
	val int
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	_, sc, k := readData(reader)

	idMap := make(map[string]int)

	for sc.Scan() {
		id := sc.Text()
		idMap[id] = idMap[id] + 1
	}

	var idSlice []KV
	for key, v := range idMap {
		idSlice = append(idSlice, KV{key, v})
	}

	sort.Slice(idSlice, func(i, j int) bool {
		if idSlice[i].val == idSlice[j].val {
			return idSlice[i].key < idSlice[j].key
		}

		return idSlice[i].val > idSlice[j].val
	})

	var b strings.Builder
	for i := 0; i < k; i++ {
		b.WriteString(idSlice[i].key + " ")
	}

	result := strings.TrimRight(b.String(), " ")

	writer.WriteString(result)
	writer.WriteString("\n")

	writer.Flush()
}

func readData(reader *bufio.Reader) (n int, sc *bufio.Scanner, k int) {
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line3, _ := reader.ReadString('\n')

	n, _ = strconv.Atoi(strings.TrimRight(line1, "\n"))
	sc = createWordsScanner(line2)
	k, _ = strconv.Atoi(strings.TrimRight(line3, "\n"))

	return
}

func createWordsScanner(line string) *bufio.Scanner {
	strReader := strings.NewReader(line)
	strBufReader := bufio.NewReader(strReader)
	scanner := bufio.NewScanner(strBufReader)
	scanner.Split(bufio.ScanWords)

	return scanner
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
