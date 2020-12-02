package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type IOOperator struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

func main() {
	file := openFile("input.txt")
	defer file.Close()

	ioOperator := createIOOperator(file, os.Stdout)

	ioOperator.ProcessHeadSegment()
	homeCounter := ioOperator.ProcessBodySegment()
	ioOperator.ProcessTailSegment(homeCounter)

	ioOperator.writer.WriteString("\n")
	ioOperator.writer.Flush()

	fmt.Println("0 0 1 0 0 0 0 1 0 0 1 2 1 0 1 0 1 0 0 0") // TODO: remove
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}

func createIOOperator(input *os.File, output *os.File) IOOperator {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	scanner.Scan() // read useless n

	writer := bufio.NewWriter(output)

	return IOOperator{scanner, writer}
}

func (ioOperator IOOperator) WriteStringWithSpace(str string) {
	ioOperator.writer.WriteString(str)
	ioOperator.writer.WriteString(" ")
}

func (ioOperator IOOperator) ProcessHeadSegment() {
	writeHeadSegment := func(homeCounter int) {
		for i := homeCounter; i > 0; i-- {
			ioOperator.WriteStringWithSpace(strconv.Itoa(i))
		}
	}

	homeCounter := 0
	for ioOperator.scanner.Scan() {
		home := ioOperator.scanner.Text()

		if home == "0" {
			if homeCounter > 0 {
				writeHeadSegment(homeCounter)
			}

			ioOperator.WriteStringWithSpace(home)
			break
		} else {
			homeCounter++
		}
	}
}

func (ioOperator IOOperator) ProcessBodySegment() int {
	writeBodySegment := func(homeCounter int) {
		halfOfSegment := homeCounter / 2
		for i := 1; i <= halfOfSegment+homeCounter%2; i++ {
			ioOperator.WriteStringWithSpace(strconv.Itoa(i))
		}
		for i := halfOfSegment; i > 0; i-- {
			ioOperator.WriteStringWithSpace(strconv.Itoa(i))
		}
	}

	homeCounter := 0
	for ioOperator.scanner.Scan() {
		home := ioOperator.scanner.Text()

		if home == "0" {
			if homeCounter > 0 {
				writeBodySegment(homeCounter)
				homeCounter = 0
			}

			ioOperator.WriteStringWithSpace(home)
		} else {
			homeCounter++
		}
	}

	return homeCounter
}

func (ioOperator IOOperator) ProcessTailSegment(homeCounter int) {
	if homeCounter == 0 {
		return
	}

	writeTailSegment := func(homeCounter int) {
		for i := 1; i <= homeCounter; i++ {
			ioOperator.writer.WriteString(strconv.Itoa(i))
			if i < homeCounter {
				ioOperator.writer.WriteString(" ")
			}
		}
	}

	writeTailSegment(homeCounter)
}
