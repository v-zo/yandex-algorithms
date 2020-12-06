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

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	scanner.Scan()
	now, _ := strconv.Atoi(scanner.Text())
	prev := -273
	var counter int
	for scanner.Scan() {
		next, _ := strconv.Atoi(scanner.Text())
		if prev < now && now > next {
			counter++
		}

		prev = now
		now = next
	}

	if prev < now {
		counter++
	}

	fmt.Println(counter)
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
