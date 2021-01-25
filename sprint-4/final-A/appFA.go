/*

посылка https://contest.yandex.ru/contest/24414/problems/A/

-- ПРИНЦИП РАБОТЫ --
	Сначало строим поисковый индекс - мап, который сопоставляет словам массив индексов документов, в которые данное
слово входит. Затем считываем запросы и с помощью поискового индекса находим релевантность.
	Для хранения релевантности используется мап relMap. В нем по индексу документа хранится счетчик
вхождения в документ текущей строки. Далее Преобразуем мап в слайс чтобы можно было применть сортировку.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
	По сути мы просто перебираем все слова во всех документах и считаем количество вхождений. Мы предварительно
"кешируем" результаты поиска в индексе, что является лишь оптимизацией и в целом алгоритм делает именно то что
нужно сделать согласно заданию.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
	При построении индекса мы пробегаем по всем словам всех документов. Сложность O(D*W), где D - общее количество
документов, W - среднее (характерное) кол-во слов в документе. Эта процедура дорогая, но выполняется 1 раз, поэтому
при большом количестве запросов окупается пропорционально числу запросов.
	Сложность вычисления релевантности O(R+D*W), где R - количество запросов. Видно, что при малых R сложность
определяется числом документов и "средней загруженностью" документа словами. При R -> ∞ сложность алгоритма
линейна по R.

*/

package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

const maxDocsPerLine = 5

func main() {
	file := openFile("input.txt")
	defer file.close()

	reader := bufio.NewReader(file.osFile)
	writer := bufio.NewWriter(os.Stdout)

	Solve(reader, writer)
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	yaReader := &YaReader{reader}
	docs, queries := readData(yaReader)

	rsw := &RelevanceSliceWriter{writer}
	si := buildSearchIndex(docs)
	for _, query := range queries {
		relevanceSlice := si.queryRelevanceSlice(query)
		rsw.writeRelevanceSlice(relevanceSlice)
	}

	err := writer.Flush()
	check(err)
}

type DocumentRelevance struct {
	docId int
	count int
}

type RelevanceSliceWriter struct {
	writer *bufio.Writer
}

func (rsw *RelevanceSliceWriter) writeCarefully(s string) {
	_, err := rsw.writer.WriteString(s)
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
	words := uniqueWords(strings.Fields(query))
	relevanceSlice := si.getRelevanceSlice(words)
	sortRelSlice(relevanceSlice)

	return relevanceSlice
}

func uniqueWords(words []string) (uw []string) {
	set := make(map[string]struct{})
	for _, word := range words {
		set[word] = struct{}{}
	}

	for word := range set {
		uw = append(uw, word)
	}

	return
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
	index map[string][]int
}

func (si *SearchIndex) getRelevanceSlice(words []string) []DocumentRelevance {
	relevanceMap := make(map[int]int)
	for _, word := range words {
		includedInDocs := si.index[word]
		for _, doc := range includedInDocs {
			relevanceMap[doc] = relevanceMap[doc] + 1
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

func buildSearchIndex(docs []string) *SearchIndex {
	searchIndex := make(map[string][]int)

	for i, doc := range docs {
		words := strings.Fields(doc)
		for _, word := range words {
			searchIndex[word] = append(searchIndex[word], i)
		}
	}

	return &SearchIndex{searchIndex}
}

func readData(reader *YaReader) (docs []string, queries []string) {
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
	line, err := reader.ReadString('\n')
	check(err)
	return strings.TrimRight(line, "\n")
}

func (reader *YaReader) readInt() int {
	line, err := reader.ReadString('\n')
	check(err)
	res, err := strconv.Atoi(strings.TrimRight(line, "\n"))
	check(err)
	return res
}

type File struct {
	osFile *os.File
}

func openFile(path string) *File {
	osFile, err := os.Open(path)
	check(err)

	return &File{osFile}
}

func (f *File) close() {
	err := f.osFile.Close()
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
