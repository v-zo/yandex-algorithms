package main

import (
	"bufio"
	"os"
)

func main() {
	file := openFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	//n, _ := strconv.Atoi(scanner.Text())

	writer := bufio.NewWriter(os.Stdout)

	writeHome := func(home string) {
		writer.WriteString(home)
		writer.WriteString(" ")
	}

	writeBodySegment := func(homeCounter int) {
		for i := 0; i < homeCounter; i++ {
			writeHome("*")
		}
	}

	writeHeadSegment := func(homeCounter int) {
		for i := 0; i < homeCounter; i++ {
			writeHome("x")
		}
	}

	writeTailSegment := func(homeCounter int) {
		for i := 0; i < homeCounter; i++ {
			writer.WriteString("*")
			if i < homeCounter-1 {
				writer.WriteString(" ")
			}
		}
	}

	homeCounter := 0
	isHead := true
	for scanner.Scan() {
		home := scanner.Text()

		if home == "0" {
			if homeCounter > 0 {
				if isHead {
					writeHeadSegment(homeCounter)
				} else {
					writeBodySegment(homeCounter)
				}

				isHead = false
				homeCounter = 0
			}

			writeHome(home)
		} else {
			homeCounter++
		}
	}

	if homeCounter > 0 {
		writeTailSegment(homeCounter)
	}

	//writer.WriteString("\n")
	writer.Flush()

	//fmt.Println("\n")
	//fmt.Println("2 1 0 1 2 1 0 1 0 0 1 0 0 1 0 1 2 1 0 1")
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
