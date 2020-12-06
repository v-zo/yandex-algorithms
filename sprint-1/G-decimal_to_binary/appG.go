package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file := openFile("input.txt")
	defer file.Close()

	dec := readDecimal(file)

	var bin string
	quotient := dec
	for quotient > 1 {
		bin = strconv.Itoa(quotient%2) + bin
		quotient = quotient / 2
	}

	if quotient == 1 {
		bin = "1" + bin
	}

	fmt.Println(bin)
}

func readDecimal(input *os.File) int {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	str := scanner.Text()
	num, _ := strconv.Atoi(str)

	return num
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
