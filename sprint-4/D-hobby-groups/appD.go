package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file := openFile("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	Solve(reader)
}

type Entry struct {
	k string
	v int
}

func Solve(reader *bufio.Reader) {
	yaReader := &YaReader{reader}
	n := yaReader.readInt()

	log := make(map[string]int)

	for i := 1; i <= n; i++ {
		s := yaReader.readString()
		if log[s] == 0 {
			log[s] = i
		}
	}

	var arr []Entry
	for k, v := range log {
		arr = append(arr, Entry{k, v})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].v < arr[j].v
	})

	for _, a := range arr {
		fmt.Println(a.k)
	}
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
