/*

посылка
https://contest.yandex.ru/contest/24414/run-report/47293035/

-- ПРИНЦИП РАБОТЫ --
	Релевантность документа для запроса - это сумма вхождений в данный документ каждого слова этого запроса.
Если таких запросов несколько, то для каждого такого запроса мы будем итерироваться по словам одних и тех же
документов многократно.
	Идея алгоритма в том, чтобы единожды запомнить результаты сканирования документов.
Это удобно сделать предварительно построив поисковый индекс - объект, который хранит информацию в разрезе слов,
а не документов. Тогда при каждом запросе мы имеем дело уже с одним и тем же набором данных - отображением коллекции
слов на списки документов в которые эти слова входят.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
	После построения поискового индекса для каждого слова запроса нам известен список документиов
в которых оно содержится. На основании этого легко посчитать искомую релевантность каждого документа,
посчитав суммарное количество слов, которые входят в данный документ.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
	При построении индекса мы перебираем все слова всех документов. Сложность O(D*W), где D - общее количество
документов, W - среднее (характерное) кол-во слов в документе. Эта процедура дорогая, но выполняется 1 раз, поэтому
при большом количестве запросов окупается пропорционально числу запросов.
	Сложность вычисления релевантности O(R+D*W), где R - количество запросов. Видно, что при малых R сложность
определяется числом документов и "средней загруженностью" документа словами. При R -> ∞ сложность алгоритма
линейна по R.

*/

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
	yaReader := &YaReader{reader}
	docs, queries := readData(yaReader)

	rsw := &RelevanceSliceWriter{writer}
	si := buildSearchIndex(docs)
	for _, query := range queries {
		relevanceSlice := si.queryRelevanceSlice(query)
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
