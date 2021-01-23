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
		words := uniqueWords(strings.Fields(query))

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

		sortRelSlice(relSlice)

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

func sortRelSlice(r []Rel) {
	sort.Slice(r, func(i, j int) bool {
		if r[i].count == r[j].count {
			return r[i].doc < r[j].doc
		}

		return r[i].count > r[j].count
	})
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
