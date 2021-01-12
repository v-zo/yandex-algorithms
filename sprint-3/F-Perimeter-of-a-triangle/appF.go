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
	_, nums := readData(reader)

	res := findSol(nums)

	writer.WriteString(strconv.Itoa(res))
	writer.WriteString("\n")
	writer.Flush()
}

func findSol(nums []int) int {
	n := len(nums)
	SortIntReverse(nums)

	if nums[0] < nums[1]+nums[2] {
		return nums[0] + nums[1] + nums[2]
	}

	f := -1
	lo := 0
	hi := n - 1
	i := hi / 2
	for hi > lo && i > 0 {
		if nums[i] < nums[i+1]+nums[i+2] {
			hi = i
		} else {
			lo = i
			f = i + 1
		}

		f = i

		i = (hi + lo) / 2
		if i > n-3 {
			break
		}
	}

	return nums[f] + nums[f+1] + nums[f+2]
}

func SortIntReverse(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
}

func readData(reader *bufio.Reader) (n int, nums []int) {
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')

	n, _ = strconv.Atoi(strings.Trim(line1, "\n"))
	numsStr := strings.Fields(strings.Trim(line2, "\n"))

	nums = strToInt(numsStr)

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
