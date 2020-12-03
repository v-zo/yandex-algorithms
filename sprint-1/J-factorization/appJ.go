package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Array []int
type PrimeFactors []int

func main() {
	file := openFile("input.txt")
	n := readData(file)
	file.Close()

	var pf PrimeFactors
	pf = pf.FindFactors(n)

	fmt.Println(pf)
}

func (pf PrimeFactors) FindFactors(n int) PrimeFactors {
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			pf = append(pf, i)
			return pf.FindFactors(n / i)
		}
	}

	return pf
}

func GetPrimes(n int) []int {
	var numbers Array = []int{0, 0}
	for i := 2; i <= n; i++ {
		numbers = append(numbers, i)
	}

	for num := 2; num <= n; num++ {
		if numbers[num] > 0 {
			for j := 2 * num; j <= n; j += num {
				numbers[j] = 0
			}
		}
	}

	return numbers.Filter(testNonZero)
}

func testNonZero(item int) bool {
	if item > 0 {
		return true
	}

	return false
}

func (arr Array) Filter(test func(int) bool) []int {
	var filteredArr []int

	for _, item := range arr {
		if test(item) {
			filteredArr = append(filteredArr, item)
		}
	}

	return filteredArr
}

func readData(inputFile *os.File) int {
	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	return n
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
