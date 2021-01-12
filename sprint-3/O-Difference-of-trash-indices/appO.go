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
	_, nums, k := readData(reader)

	res := findSol(nums, k)

	var _ = res
	writer.WriteString(strconv.Itoa(res))
	writer.WriteString("\n")
	writer.Flush()
}

func findSol(nums []int, k int) (res int) {
	SortInt(nums)

	depth := 1
	N := len(nums) - depth

	for k > N {
		depth++
		k = k - N
		N = N - 1
	}

	res = mapToDeltas(nums, depth)[k-1]

	return
}

func mapToDeltas(nums []int, depth int) (res []int) {
	N := len(nums) - depth
	res = make([]int, N)
	for i := 0; i < N; i++ {
		res[i] = abs(nums[i+depth] - nums[i])
	}

	SortInt(res)

	return
}

func SortInt(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func readData(reader *bufio.Reader) (n int, nums []int, k int) {
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line3, _ := reader.ReadString('\n')

	n, _ = strconv.Atoi(strings.Trim(line1, "\n"))
	numsStr := strings.Fields(strings.Trim(line2, "\n"))
	k, _ = strconv.Atoi(strings.Trim(line3, "\n"))

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
