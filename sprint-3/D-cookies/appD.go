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
	n, children, m, cookies := readData(reader)

	SortInt(children)
	SortInt(cookies)

	if n > m {
		children = children[:m]
	}

	ch := 0
	for i := 0; i < m && ch < n; i++ {
		if cookies[i] >= children[ch] {
			ch++
		}
	}

	writer.WriteString(strconv.Itoa(ch))
	writer.WriteString("\n")
	writer.Flush()
}

func SortInt(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func readData(reader *bufio.Reader) (n int, children []int, m int, cookies []int) {
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line3, _ := reader.ReadString('\n')
	line4, _ := reader.ReadString('\n')

	n, _ = strconv.Atoi(strings.Trim(line1, "\n"))
	childrenStr := strings.Fields(strings.Trim(line2, "\n"))
	m, _ = strconv.Atoi(strings.Trim(line3, "\n"))
	cookiesStr := strings.Fields(strings.Trim(line4, "\n"))

	children = strToInt(childrenStr)
	cookies = strToInt(cookiesStr)

	return
}

func strToInt(arr []string) (integers []int) {
	integers = []int{}
	for _, s := range arr {
		i, _ := strconv.Atoi(s)
		integers = append(integers, i)
	}

	return
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
