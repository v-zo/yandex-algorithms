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
	scanner.Scan() // read useless n

	writer := bufio.NewWriter(os.Stdout)

	writeStringWithSpace := func(str string) {
		writer.WriteString(str)
		writer.WriteString(" ")
	}

	writeHeadSegment := func(homeCounter int) {
		for i := homeCounter; i > 0; i-- {
			writeStringWithSpace(strconv.Itoa(i))
		}
	}

	writeBodySegment := func(homeCounter int) {
		halfOfSegment := homeCounter / 2
		for i := 1; i <= halfOfSegment+homeCounter%2; i++ {
			writeStringWithSpace(strconv.Itoa(i))
		}
		for i := halfOfSegment; i > 0; i-- {
			writeStringWithSpace(strconv.Itoa(i))
		}
	}

	writeTailSegment := func(homeCounter int) {
		for i := 1; i <= homeCounter; i++ {
			writer.WriteString(strconv.Itoa(i))
			if i < homeCounter {
				writer.WriteString(" ")
			}
		}
	}

	homeCounter := 0
	for scanner.Scan() {
		home := scanner.Text()

		if home == "0" {
			if homeCounter > 0 {
				writeHeadSegment(homeCounter)
			}

			writeStringWithSpace(home)
			break
		} else {
			homeCounter++
		}
	}

	homeCounter = 0
	for scanner.Scan() {
		home := scanner.Text()

		if home == "0" {
			if homeCounter > 0 {
				writeBodySegment(homeCounter)
				homeCounter = 0
			}

			writeStringWithSpace(home)
		} else {
			homeCounter++
		}
	}

	if homeCounter > 0 {
		writeTailSegment(homeCounter)
	}

	writer.WriteString("\n")
	writer.Flush()

	fmt.Println("0 0 1 0 0 0 0 1 0 0 1 2 1 0 1 0 1 0 0 0") // TODO: remove
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
