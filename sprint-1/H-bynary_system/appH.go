package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file := openFile("input.txt")
	defer file.Close()

	a, b := readTerms(file)
	a, b = prependZeroes(a, b)

	result := getSum(a, b)

	fmt.Println(result)

}

func getSum(a string, b string) string {
	var result string
	var prepend string
	var surplus int
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '0' && b[i] == '0' {
			if surplus > 0 {
				prepend = "1"
				surplus--
			} else {
				prepend = "0"
			}
		}

		if a[i] != b[i] {
			if surplus > 0 {
				prepend = "0"
			} else {
				prepend = "1"
			}
		}

		if a[i] == '1' && b[i] == '1' {
			if surplus > 0 {
				prepend = "1"

			} else {
				prepend = "0"
				surplus++
			}

		}

		result = prepend + result
	}

	return strings.Repeat("1", surplus) + result
}

func prependZeroes(a string, b string) (string, string) {
	diff := len(a) - len(b)
	if diff > 0 {
		return a, strings.Repeat("0", diff) + b
	}

	if diff < 0 {
		return strings.Repeat("0", -diff) + a, b
	}

	return a, b
}

func readTerms(input *os.File) (string, string) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	a := scanner.Text()
	scanner.Scan()
	b := scanner.Text()

	return a, b
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
