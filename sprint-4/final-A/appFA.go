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

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	yaReader := &YaReader{reader}
	docs, queries := readData(yaReader)

	si := buildSearchIndex(docs)

	findDocs(queries, si, writer)

	writer.Flush()
}

type Rel struct {
	doc   int
	count int
}

func findDocs(queries []string, si map[string][]int, writer *bufio.Writer) {
	for _, query := range queries {
		words := strings.Fields(query)

		relMap := make(map[int]int)
		for _, word := range words {
			includedInDocs := si[word]
			if len(includedInDocs) > 0 {
				for _, doc := range includedInDocs {
					relMap[doc] = relMap[doc] + 1
				}
			}
		}

		var relSlice []Rel
		for doc, count := range relMap {
			relSlice = append(relSlice, Rel{doc, count})
		}

		sort.Slice(relSlice, func(i, j int) bool {
			if relSlice[i].count == relSlice[j].count {
				return relSlice[i].doc < relSlice[j].doc
			}

			return relSlice[i].count > relSlice[j].count
		})

		imax := 5
		if len(relSlice) < 5 {
			imax = len(relSlice)
		}

		if imax > 0 {
			for i := 0; i < imax; i++ {
				ch := strconv.Itoa(relSlice[i].doc + 1)
				writer.WriteString(ch)
				if i != imax-1 {
					writer.WriteString(" ")
				}
			}

			writer.WriteString("\n")
		}
	}
}

func buildSearchIndex(docs []string) map[string][]int {
	searchIndex := make(map[string][]int)

	for i, doc := range docs {
		words := strings.Fields(doc)
		for _, word := range words {
			searchIndex[word] = append(searchIndex[word], i)
		}
	}

	return searchIndex
}

func readData(reader *YaReader) (docs []string, queries []string) {
	docs = []string{}
	a := reader.readInt()
	for i := 0; i < a; i++ {
		docs = append(docs, reader.readString())
	}
	m := reader.readInt()
	for i := 0; i < m; i++ {
		queries = append(queries, reader.readString())
	}

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
