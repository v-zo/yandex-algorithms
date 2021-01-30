package main

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const maxDocsPerLine = 5

func main() {
	file := openFile("input.txt")
	defer file.close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	Solve(reader, writer)

	err := writer.Flush()
	check(err)
}

func Solve(reader *bufio.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)

	rsw := &RelevanceSliceWriter{writer}
	si := buildSearchIndex(scanner)
	scanner.Scan()
	maxQueries, err := strconv.Atoi(scanner.Text())
	check(err)

	for i := 0; i < maxQueries; i++ {
		scanner.Scan()
		relevanceSlice := si.queryRelevanceSlice(scanner.Text())
		rsw.writeRelevanceSlice(relevanceSlice)
	}
}

type DocumentRelevance struct {
	docId int
	count int
}

type RelevanceSliceWriter struct {
	writer io.Writer
}

func (rsw *RelevanceSliceWriter) writeCarefully(s string) {
	_, err := io.WriteString(rsw.writer, s)
	check(err)
}

func (rsw *RelevanceSliceWriter) writeRelevanceSlice(rs []DocumentRelevance) {
	maxDocs := maxDocsPerLine
	if len(rs) < maxDocsPerLine {
		maxDocs = len(rs)
	}

	stringifyDocId := func(i int) string {
		return strconv.Itoa(rs[i].docId + 1)
	}

	if maxDocs > 0 {
		docId := stringifyDocId(0)
		rsw.writeCarefully(docId)
		for i := 1; i < maxDocs; i++ {
			rsw.writeCarefully(" ")
			docId = stringifyDocId(i)
			rsw.writeCarefully(docId)
		}

		rsw.writeCarefully("\n")
	}
}

func (si *SearchIndex) queryRelevanceSlice(query string) []DocumentRelevance {
	relevanceSlice := si.getRelevanceSlice(query)
	sortRelSlice(relevanceSlice)

	return relevanceSlice
}

func sortRelSlice(r []DocumentRelevance) {
	sort.Slice(r, func(i, j int) bool {
		if r[i].count == r[j].count {
			return r[i].docId < r[j].docId
		}

		return r[i].count > r[j].count
	})
}

type SearchIndex struct {
	index map[string]map[int]int
}

func (si *SearchIndex) getRelevanceSlice(query string) []DocumentRelevance {
	stringReader := strings.NewReader(query)
	sc := bufio.NewScanner(stringReader)
	sc.Split(bufio.ScanWords)
	relevanceMap := make(map[int]int)
	uniqueWords := make(map[string]struct{})

	for sc.Scan() {
		word := sc.Text()
		if _, has := uniqueWords[word]; !has {
			uniqueWords[word] = struct{}{}
			includedInDocs := si.index[word]
			for doc, count := range includedInDocs {
				relevanceMap[doc] = relevanceMap[doc] + count
			}
		}
	}

	return mapToRelevanceSlice(relevanceMap)
}

func mapToRelevanceSlice(m map[int]int) (relevanceSlice []DocumentRelevance) {
	for doc, count := range m {
		relevanceSlice = append(relevanceSlice, DocumentRelevance{doc, count})
	}

	return
}

func buildSearchIndex(scanner *bufio.Scanner) *SearchIndex {
	scanner.Scan()
	maxDocs, err := strconv.Atoi(scanner.Text())
	check(err)

	searchIndex := make(map[string]map[int]int)

	for i := 0; i < maxDocs; i++ {
		scanner.Scan()
		words := strings.Fields(scanner.Text())
		for _, word := range words {
			sw := searchIndex[word]
			if sw == nil {
				searchIndex[word] = make(map[int]int)
			}
			searchIndex[word][i] = sw[i] + 1
		}
	}

	return &SearchIndex{searchIndex}
}

type File struct {
	*os.File
}

func openFile(path string) *File {
	osFile, err := os.Open(path)
	check(err)

	return &File{osFile}
}

func (file *File) close() {
	err := file.Close()
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
