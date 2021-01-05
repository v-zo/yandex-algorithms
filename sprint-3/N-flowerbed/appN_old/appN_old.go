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
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	scanner.Scan()

	readFb := func() (start int, end int) {
		line := scanner.Text()
		fields := strings.Fields(line)
		start, _ = strconv.Atoi(fields[0])
		end, _ = strconv.Atoi(fields[1])

		return
	}

	var fbs [][2]int
	seedStart, seedEnd := readFb()
	fbs = append(fbs, [2]int{seedStart, seedEnd})

	for scanner.Scan() {
		start, end := readFb()
		merged := mergeFb(fbs, start, end)

		if !merged {
			fbs = append(fbs, [2]int{start, end})
		}
	}

	sort.Slice(fbs, func(i, j int) bool {
		return fbs[i][0] < fbs[j][0]
	})

	for _, fb := range fbs {
		writer.WriteString(strconv.Itoa(fb[0]) + " " + strconv.Itoa(fb[1]))
		writer.WriteString("\n")
	}

	writer.Flush()
}

func mergeFb(fbs [][2]int, start int, end int) (merged bool) {
	merged = false

	for i, pair := range fbs {
		L := pair[0]
		R := pair[1]

		if start <= L && end >= L {
			fbs[i][0] = start
			if end >= R {
				fbs[i][1] = end
			} else {
				fbs[i][1] = R
			}

			merged = true
		}
		if start >= L && start <= R {
			fbs[i][0] = L
			if end >= R {
				fbs[i][1] = end
			} else {
				fbs[i][1] = R
			}

			merged = true
		}
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
