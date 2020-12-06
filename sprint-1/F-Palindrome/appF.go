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

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	var stringBuilder strings.Builder
	for scanner.Scan() {
		char := scanner.Text()

		if isLetterOrNum(char) {
			stringBuilder.WriteString(strings.ToLower(char))
		}
	}

	cleanString := stringBuilder.String()

	result := "True"
	length := len(cleanString)
	for i := 0; i < length; i++ {
		if cleanString[i] != cleanString[length-i-1] {
			result = "False"
		}
	}

	fmt.Println(result)
}

func isLetterOrNum(c string) bool {
	return ("a" <= c && c <= "z") || ("A" <= c && c <= "Z") || ("0" <= c && c <= "9")
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
